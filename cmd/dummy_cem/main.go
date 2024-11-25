// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/service"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	eglpc "github.com/enbility/eebus-go/usecases/eg/lpc"
	"github.com/enbility/eebus-grpc/utils/certificates"
	log "github.com/enbility/eebus-grpc/utils/logging"
	shipapi "github.com/enbility/ship-go/api"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
)

var remoteSki string

type hems struct {
	myService *service.Service

	uceglpc ucapi.EgLPCInterface

	log.Logger
}

func (h *hems) run(
	eebusPort int,
	remoteSki string,
	certificateFilePath string,
	privateKeyFilePath string,
) {
	certificate := certificates.SetupCertificate(
		certificateFilePath,
		privateKeyFilePath,
		"Dummy",
		"Dummy",
		"DE",
		"Dummy-Unit-01",
	)

	configuration, err := api.NewConfiguration(
		"Demo", "Demo", "HEMS", "123456789",
		[]shipapi.DeviceCategoryType{shipapi.DeviceCategoryTypeEnergyManagementSystem},
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeCEM},
		eebusPort, certificate, time.Second*4)
	if err != nil {
		h.Errorf("Error creating configuration: %s", err)
		os.Exit(1)
	}
	configuration.SetAlternateIdentifier("Dummy-Unit-01")

	h.myService = service.NewService(configuration, h)
	h.myService.SetLogging(h)

	if err = h.myService.Setup(); err != nil {
		h.Errorf("Error setting up service: %s", err)
		os.Exit(1)
	}

	localEntity := h.myService.LocalDevice().EntityForType(model.EntityTypeTypeCEM)
	h.uceglpc = eglpc.NewLPC(localEntity, nil)
	h.myService.AddUseCase(h.uceglpc)

	h.myService.RegisterRemoteSKI(remoteSki)

	h.myService.Start()
	// defer h.myService.Shutdown()
}

func (h *hems) sendDummyLoadLimits() {
	// Print currently connected remote services
	for _, remoteDevice := range h.myService.LocalDevice().RemoteDevices() {
		h.Infof("found device: %v", *remoteDevice.DeviceType())
		for _, entity := range remoteDevice.Entities() {
			if entity.EntityType() == model.EntityTypeTypeDeviceInformation {
				continue
			}
			h.Infof("found entity: %v", entity.EntityType())
			scenarios := h.uceglpc.AvailableScenariosForEntity(entity)
			if len(scenarios) == 0 {
				h.Errorf("0 scenarios")
			}
			if !slices.Contains(scenarios, 1) {
				h.Errorf("scenario 1 not supproted")
			}
		}
	}

	remoteEntityScenarios := h.uceglpc.RemoteEntitiesScenarios()
	if len(remoteEntityScenarios) == 0 {
		h.Infof("No remote entities connected")
		return
	}
	h.Infof("remote entities connected..")
	var entities []spineapi.EntityRemoteInterface
	for _, remoteEntityScenario := range remoteEntityScenarios {
		ski := remoteEntityScenario.Entity.Device().Ski()
		addressString := remoteEntityScenario.Entity.Address().String()
		h.Infof("Remote entity: %s, %s", ski, addressString)
		for _, scenario := range remoteEntityScenario.Scenarios {
			h.Infof("Scenario: %d", scenario)
		}
		entities = append(entities, remoteEntityScenario.Entity)
	}

	// random value between 100 and 3000
	limit := 100 + (3000-100)*rand.Float64()
	// Send dummy load limits
	for _, entity := range entities {
		remoteEntity := entity
		loadLimit := ucapi.LoadLimit{
			Duration:     time.Second * 5,
			IsChangeable: true,
			IsActive:     true,
			Value:        limit,
		}
		h.uceglpc.WriteConsumptionLimit(
			remoteEntity,
			loadLimit,
			func(result model.ResultDataType) {
				if result.ErrorNumber != nil {
					h.Errorf("Result Error: %d", *result.ErrorNumber)
				}
				if result.Description != nil {
					h.Infof("Result: %s", *result.Description)
				}
			},
		)
	}
}

func (h *hems) DummyLimitsRoutine() {
	h.Infof("Start limits routine..\n")
	for {
		h.sendDummyLoadLimits()
		time.Sleep(5 * time.Second)
	}
}

// EEBUSServiceHandler

func (h *hems) RemoteSKIConnected(service api.ServiceInterface, ski string) {
	fmt.Println("remote ski connected: ", ski)
}

func (h *hems) RemoteSKIDisconnected(service api.ServiceInterface, ski string) {
	fmt.Println("remote ski disconnected: ", ski)
}

func (h *hems) VisibleRemoteServicesUpdated(service api.ServiceInterface, entries []shipapi.RemoteService) {
}

func (h *hems) ServiceShipIDUpdate(ski string, shipdID string) {}

func (h *hems) ServicePairingDetailUpdate(ski string, detail *shipapi.ConnectionStateDetail) {
	if ski == remoteSki && detail.State() == shipapi.ConnectionStateRemoteDeniedTrust {
		fmt.Println("The remote service denied trust. Exiting.")
		h.myService.CancelPairingWithSKI(ski)
		h.myService.UnregisterRemoteSKI(ski)
		h.myService.Shutdown()
		os.Exit(0)
	}
}

func (h *hems) AllowWaitingForTrust(ski string) bool {
	return ski == remoteSki
}

func main() {
	var eebusPort = flag.Int("eebus-port", 1234, "The port to listen on for EEBUS connections")
	var remoteSki = flag.String("remote-ski", "", "The remote SKI to connect to")
	var certificateFilePath = flag.String("certificate-path", "", "The path to the certificate file")
	var privateKeyFilePath = flag.String("private-key-path", "", "The path to the private key file")

	flag.Parse()

	if *remoteSki == "" {
		log.Info("No remote ski provided")
	}
	if *certificateFilePath == "" {
		log.Error("No certificate path provided")
		os.Exit(1)
	}
	if *privateKeyFilePath == "" {
		log.Error("No private key path provided")
		os.Exit(1)
	}
	log.Infof("Certificate path: %s", *certificateFilePath)
	log.Infof("Private key path: %s", *privateKeyFilePath)
	log.Infof("Remote ski: %s", *remoteSki)
	log.Infof("EEBUS port: %d", *eebusPort)

	h := hems{}
	h.run(
		*eebusPort,
		*remoteSki,
		*certificateFilePath,
		*privateKeyFilePath,
	)

	go h.DummyLimitsRoutine()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit
}
