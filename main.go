package main

import (
	slasherConfig "github.com/OCRVblockchain/slasher/pkg/config"
	"github.com/OCRVblockchain/slasher/pkg/core"
	"github.com/OCRVblockchain/slasher/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf, err := slasherConfig.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	slasher, err := core.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = slasher.EnrollUser(conf.ActionUser, conf.Secret)
	if err != nil {
		log.Fatal(err)
	}

	if conf.Mode == "revokecert" || conf.Mode == "fullslash" {

		revocationResponse, err := slasher.Revoke()
		if err != nil {
			log.Fatal(err)
		}
		helpers.ShowRevoked(revocationResponse)

	}
	if conf.Mode == "removeidentity" || conf.Mode == "fullslash" {

		identityResponse, err := slasher.RemoveIdentity()
		if err != nil {
			log.Fatal(err)
		}
		helpers.ShowRemoved(identityResponse)

	}
}
