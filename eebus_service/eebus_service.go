// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package eebus_service

import (
	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/service"
	shipapi "github.com/enbility/ship-go/api"
)

type Service struct {
	service.Service
	isConnected bool
}

func NewService() *Service {
	return &Service{
		isConnected: false,
	}
}

func (h *Service) RemoteSKIConnected(service api.ServiceInterface, ski string) {
	h.isConnected = true
}

func (h *Service) RemoteSKIDisconnected(service api.ServiceInterface, ski string) {
	h.isConnected = false
}

func (h *Service) VisibleRemoteServicesUpdated(service api.ServiceInterface, entries []shipapi.RemoteService) {
}

func (h *Service) ServiceShipIDUpdate(ski string, shipdID string) {
}

func (h *Service) ServicePairingDetailUpdate(ski string, detail *shipapi.ConnectionStateDetail) {
}

func (h *Service) AllowWaitingForTrust(ski string) bool {
	return true
}
