package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"
	"github.com/mileusna/crontab"
)

const BitSize int = 2048
const PathPem string = "./pem/"

func OpenPemFactory() {

	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, BitSize)
	checkError(err)

	publicKey := key.PublicKey
	makePrivateKey(key)
	makePublicKey(publicKey)
	schedulerForKeysCreation()
}

func schedulerForKeysCreation() {
	timeScheduler := os.Getenv("SCHEDULER_FOR_KEYS_PEM_CREATE")
	if len(timeScheduler) > 0 {
		ctab := crontab.New() // create cron table
		ctab.MustAddJob(timeScheduler, OpenPemFactory)
	}
}

func makePrivateKey(key *rsa.PrivateKey) {
	outFile, err := os.Create(PathPem + "private.pem")
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func makePublicKey(pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(PathPem + "public.pem")
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
