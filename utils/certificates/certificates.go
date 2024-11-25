// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package certificates

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"os"

	log "github.com/enbility/eebus-grpc/utils/logging"
	"github.com/enbility/ship-go/cert"
)

func CreateCertifacteFile(certificate *tls.Certificate, path string) error {
	var pemdata []byte = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate.Certificate[0],
	})

	if _, err := os.Stat(path); err == nil {
		log.Errorf("Failed to create certificate file \"%s\", because it already exists", path)
		return os.ErrExist
	}
	err := os.WriteFile(path, pemdata, 0644)
	if err != nil {
		log.Errorf("Failed to write certificate file \"%s\"", path)
		return err
	}

	return nil
}

func CreatePrivateKeyFile(certificate *tls.Certificate, path string) error {
	var keydata []byte
	var err error

	keydata, err = x509.MarshalECPrivateKey(certificate.PrivateKey.(*ecdsa.PrivateKey))
	if err != nil {
		log.Error("Failed to marshal private key")
		return err
	}

	var pemdata []byte = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keydata})

	if _, err := os.Stat(path); err == nil {
		log.Errorf("Failed to create private key file \"%s\", because it already exists", path)
		return os.ErrExist
	}
	err = os.WriteFile(path, pemdata, 0644)
	if err != nil {
		log.Errorf("Failed to write private key file \"%s\"", path)
		return err
	}

	return nil
}

func SetupCertificate(
	certificateFilePath string,
	privateKeyFilePath string,
	organizationalUnit string,
	organization string,
	country string,
	commonName string,
) tls.Certificate {
	var err error

	var certificate tls.Certificate

	var certificateFileExists, privateKeyFileExists bool

	if _, err := os.Stat(certificateFilePath); err == nil {
		certificateFileExists = true
	}
	if _, err := os.Stat(privateKeyFilePath); err == nil {
		privateKeyFileExists = true
	}

	if certificateFileExists && privateKeyFileExists {
		log.Infof("Certificate and private key already exist, skipping creation")

		certificate, err = tls.LoadX509KeyPair(certificateFilePath, privateKeyFilePath)
		if err != nil {
			log.Errorf("Failed to load certificate and private key: %v", err)
			os.Exit(1)
		}
	} else {
		log.Infof("Creating certificate and private key")

		certificate, err = cert.CreateCertificate(
			organizationalUnit,
			organization,
			country,
			commonName,
		)
		if err != nil {
			log.Errorf("Failed to create certificate: %v", err)
			os.Exit(1)
		}

		err = CreateCertifacteFile(&certificate, certificateFilePath)
		if err != nil {
			log.Errorf("Failed to create certificate file: %v", err)
		}

		err = CreatePrivateKeyFile(&certificate, privateKeyFilePath)
		if err != nil {
			log.Errorf("Failed to create private key file: %v", err)
		}
	}

	return certificate
}
