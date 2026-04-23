// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package eebus_service

import (
	"sync"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/service"
	"github.com/enbility/eebus-grpc/rpc_services/control_service"
	log "github.com/enbility/eebus-grpc/utils/logging"
	shipapi "github.com/enbility/ship-go/api"
)

// discoverySubscriberBufferSize is the per-subscriber channel buffer for
// DiscoveryEvent delivery. A slow subscriber that fills its buffer will
// have further events dropped (with a log entry) rather than blocking
// producers.
const discoverySubscriberBufferSize = 32

type discoverySubscriber struct {
	ch chan *control_service.DiscoveryEvent
}

type Service struct {
	service.Service
	isConnected bool

	// discoveryMu guards discoverySnapshot and discoverySubscribers.
	discoveryMu          sync.Mutex
	discoverySnapshot    map[string]shipapi.RemoteService
	discoverySubscribers []*discoverySubscriber
}

func NewService() *Service {
	return &Service{
		isConnected:          false,
		discoverySnapshot:    make(map[string]shipapi.RemoteService),
		discoverySubscribers: nil,
	}
}

func (h *Service) RemoteSKIConnected(service api.ServiceInterface, ski string) {
	h.isConnected = true
}

func (h *Service) RemoteSKIDisconnected(service api.ServiceInterface, ski string) {
	h.isConnected = false
}

// VisibleRemoteServicesUpdated is called by ship-go with the full current
// list of visible remote services every mDNS update. We diff against the
// last-seen snapshot and emit DISCOVERED / REMOVED events to all active
// SubscribeDiscoveries subscribers.
func (h *Service) VisibleRemoteServicesUpdated(_ api.ServiceInterface, entries []shipapi.RemoteService) {
	h.discoveryMu.Lock()

	newSnapshot := make(map[string]shipapi.RemoteService, len(entries))
	for _, entry := range entries {
		newSnapshot[entry.Ski] = entry
	}

	var events []*control_service.DiscoveryEvent

	// Detect newly discovered SKIs.
	for ski, entry := range newSnapshot {
		if _, existed := h.discoverySnapshot[ski]; !existed {
			events = append(events, h.buildDiscoveryEventLocked(control_service.DiscoveryEvent_DISCOVERED, entry))
		}
	}

	// Detect removed SKIs.
	for ski, entry := range h.discoverySnapshot {
		if _, stillPresent := newSnapshot[ski]; !stillPresent {
			events = append(events, h.buildDiscoveryEventLocked(control_service.DiscoveryEvent_REMOVED, entry))
		}
	}

	h.discoverySnapshot = newSnapshot

	// Copy the subscriber slice so we can release the lock before sending.
	subscribers := make([]*discoverySubscriber, len(h.discoverySubscribers))
	copy(subscribers, h.discoverySubscribers)

	h.discoveryMu.Unlock()

	for _, evt := range events {
		for _, sub := range subscribers {
			select {
			case sub.ch <- evt:
			default:
				log.Infof("DiscoveryEvent subscriber channel full; dropping event for SKI %s (type %s)",
					evt.GetRemoteSki(), evt.GetType().String())
			}
		}
	}
}

func (h *Service) ServiceShipIDUpdate(ski string, shipdID string) {
}

func (h *Service) ServicePairingDetailUpdate(ski string, detail *shipapi.ConnectionStateDetail) {
}

func (h *Service) AllowWaitingForTrust(ski string) bool {
	return true
}

// buildDiscoveryEventLocked constructs a DiscoveryEvent from a ship-go
// RemoteService entry. The caller must hold h.discoveryMu (the method
// itself does not touch mu-protected state, but it is invoked from
// locked sections).
func (h *Service) buildDiscoveryEventLocked(eventType control_service.DiscoveryEvent_Type, entry shipapi.RemoteService) *control_service.DiscoveryEvent {
	return &control_service.DiscoveryEvent{
		Type:           eventType,
		RemoteSki:      entry.Ski,
		ShipIdentifier: entry.Identifier,
		Brand:          entry.Brand,
		Model:          entry.Model,
		DeviceType:     entry.Type,
		Serial:         entry.Serial,
		IsTrusted:      h.skiTrusted(entry.Ski),
	}
}

// skiTrusted returns true if the sidecar has marked the given SKI as
// trusted at the moment this is called. Nil-safe: returns false if the
// SKI is not yet known to the hub.
func (h *Service) skiTrusted(ski string) bool {
	details := h.RemoteServiceForSKI(ski)
	if details == nil {
		return false
	}
	return details.Trusted()
}

// SubscribeDiscoveries registers a new subscriber for DiscoveryEvents and
// atomically returns a snapshot of the currently-visible services as
// DISCOVERED events. The snapshot and the live channel are produced under
// the same lock acquisition so that any VisibleRemoteServicesUpdated
// firing concurrently either happens entirely before the subscriber is
// registered (only the snapshot reflects it) or entirely after (only the
// channel receives it) — there are no duplicate replay events.
//
// Returns the snapshot slice, a receive-only channel for subsequent live
// events, and a cancel function that removes the subscriber and closes
// the channel. The channel is buffered; callers that fail to drain it
// promptly will have events dropped.
func (h *Service) SubscribeDiscoveries() (snapshot []*control_service.DiscoveryEvent, ch <-chan *control_service.DiscoveryEvent, cancel func()) {
	sub := &discoverySubscriber{
		ch: make(chan *control_service.DiscoveryEvent, discoverySubscriberBufferSize),
	}

	h.discoveryMu.Lock()
	h.discoverySubscribers = append(h.discoverySubscribers, sub)
	snapshot = make([]*control_service.DiscoveryEvent, 0, len(h.discoverySnapshot))
	for _, entry := range h.discoverySnapshot {
		snapshot = append(snapshot, h.buildDiscoveryEventLocked(control_service.DiscoveryEvent_DISCOVERED, entry))
	}
	h.discoveryMu.Unlock()

	// cancel removes this subscriber from the list and closes the channel.
	// It is safe to call multiple times; subsequent calls are no-ops.
	cancel = func() {
		h.discoveryMu.Lock()
		defer h.discoveryMu.Unlock()
		for i, s := range h.discoverySubscribers {
			if s == sub {
				h.discoverySubscribers = append(h.discoverySubscribers[:i], h.discoverySubscribers[i+1:]...)
				close(sub.ch)
				return
			}
		}
	}

	return snapshot, sub.ch, cancel
}

// CurrentDiscoveries returns a DISCOVERED event for each entry in the
// current snapshot. Kept as a public method for debugging / external
// callers; the gRPC handler uses SubscribeDiscoveries' atomic snapshot
// instead to avoid replay races.
func (h *Service) CurrentDiscoveries() []*control_service.DiscoveryEvent {
	h.discoveryMu.Lock()
	defer h.discoveryMu.Unlock()
	events := make([]*control_service.DiscoveryEvent, 0, len(h.discoverySnapshot))
	for _, entry := range h.discoverySnapshot {
		events = append(events, h.buildDiscoveryEventLocked(control_service.DiscoveryEvent_DISCOVERED, entry))
	}
	return events
}
