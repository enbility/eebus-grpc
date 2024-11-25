// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/enbility/eebus-grpc/utils/logging"

	"github.com/enbility/eebus-grpc/eebus_service"
	"github.com/enbility/eebus-grpc/rpc_server/control_service_server"
	"github.com/enbility/eebus-grpc/utils/certificates"
)

type app struct {
	eebusService         *eebus_service.Service
	controlServiceServer *control_service_server.Server

	log.Logger
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  go run generic-eebus/cmd/main.go -port=<port> -remoteski=<remote ski> -certificate-path=<certificate path>")
}

const (
	OrganizationalUnit = "EVerest"
	Organization       = "EVerest"
	Country            = "DE"
	CommonName         = "EVerest"
)

func main() {
	var rpcPort = flag.Int("port", 50051, "The server port")
	var certificateFilePath = flag.String("certificate-path", "", "The path to the certificate")
	var privateKeyFilePath = flag.String("private-key-path", "", "The path to the private key")
	flag.Parse()

	var app app
	if *certificateFilePath == "" {
		app.Error("No certificate path provided")
		usage()
		os.Exit(1)
	}
	if *privateKeyFilePath == "" {
		app.Error("No private key path provided")
		usage()
		os.Exit(1)
	}
	app.Infof("Certificate path: %s", *certificateFilePath)
	app.Infof("Private key path: %s", *privateKeyFilePath)
	app.Infof("Port: %d", *rpcPort)

	cert := certificates.SetupCertificate(
		*certificateFilePath,
		*privateKeyFilePath,
		OrganizationalUnit,
		Organization,
		Country,
		CommonName,
	)

	app.eebusService = eebus_service.NewService()
	app.controlServiceServer = control_service_server.NewServer(app.eebusService, &cert)

	app.controlServiceServer.Start(rpcPort)

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit

	app.controlServiceServer.Stop()
}
