// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/enbility/ship-go/cert"
)

type Parameters struct {
	abs_target_dir string
	name           string
}

func write_file(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprint(file, content)
	if err != nil {
		return err
	}
	return nil
}

func get_ski_from_cert(certificate tls.Certificate) (string, error) {
	leaf, err := x509.ParseCertificate(certificate.Certificate[0])
	if err != nil {
		return "no-ski", err
	}

	ski, err := cert.SkiFromCertificate(leaf)
	if err != nil {
		return "no-ski", err
	}

	return ski, nil
}

func create_ski_file(path string, certificate tls.Certificate) error {
	ski, err := get_ski_from_cert(certificate)
	if err != nil {
		return err
	}

	err = write_file(
		path,
		ski,
	)
	if err != nil {
		return err
	}

	fmt.Println("Created ski file at ", path)
	return nil
}

func get_cert_file_content(cert tls.Certificate) (string, error) {
	var res string

	res_bytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Certificate[0],
		},
	)
	res = strings.TrimSpace(string(res_bytes))

	return res, nil
}

func get_key_file_content(cert tls.Certificate) (string, error) {
	var res string

	res_bytes, err := x509.MarshalECPrivateKey(
		cert.PrivateKey.(*ecdsa.PrivateKey),
	)
	if err != nil {
		return res, err
	}
	res_bytes = pem.EncodeToMemory(
		&pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: res_bytes,
		},
	)

	res = strings.TrimSpace(string(res_bytes))

	return res, nil
}

func create_cert_file(path string, certificate tls.Certificate) error {
	cert_file_content, err := get_cert_file_content(certificate)
	if err != nil {
		return err
	}

	err = write_file(
		path,
		cert_file_content,
	)
	if err != nil {
		return err
	}

	fmt.Println("Created cert file at ", path)
	return nil
}

func create_key_file(path string, certificate tls.Certificate) error {
	key_file_content, err := get_key_file_content(certificate)
	if err != nil {
		return err
	}

	err = write_file(
		path,
		key_file_content,
	)
	if err != nil {
		return err
	}

	fmt.Println("Created key file at ", path)
	return nil
}

func create_files(target_dir string, name string) error {
	var err error
	var certificate tls.Certificate
	certificate, err = cert.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-01")
	if err != nil {
		return err
	}
	base_path := target_dir + "/" + name + "_"

	err = create_cert_file(base_path+"cert", certificate)
	if err != nil {
		return err
	}

	err = create_key_file(base_path+"key", certificate)
	if err != nil {
		return err
	}

	err = create_ski_file(base_path+"ski", certificate)
	if err != nil {
		return err
	}

	return nil
}

func parse_args(args []string) (Parameters, error) {
	var params Parameters
	var err error

	if len(args) != 3 {
		usage()
		return params, fmt.Errorf("There should be 2 additional args")
	}
	target_dir := args[1]
	params.abs_target_dir, err = filepath.Abs(target_dir)
	if err != nil {
		return params, err
	}
	params.name = args[2]

	return params, nil
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  go run /cmd/create_cert/main.go <target_dir> <name>")
}

func main() {
	params, err := parse_args(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	err = create_files(params.abs_target_dir, params.name)
	if err != nil {
		log.Fatal(err)
	}
}
