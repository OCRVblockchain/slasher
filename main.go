package main

import (
	"fmt"
	slasherConfig "github.com/OCRVblockchain/slasher/pkg/config"
	"github.com/OCRVblockchain/slasher/pkg/core"
	"github.com/OCRVblockchain/slasher/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf, err := slasherConfig.GetConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get Slasher config: %s\n", err))
	}

	slasher, err := core.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = slasher.EnrollUser(conf.ActionUser, conf.Secret)
	if err != nil {
		log.Fatal(err)
	}

	revocationResponse, err := slasher.Revoke()
	if err != nil {
		log.Fatal(err)
	}

	helpers.ShowRevoked(revocationResponse)
}
