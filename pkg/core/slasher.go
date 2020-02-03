package core

import (
	"errors"
	"fmt"
	slasherConfig "github.com/OCRVblockchain/slasher/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	//"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	log "github.com/sirupsen/logrus"
)

var (
	user   = "admin"
	secret = "adminpw"
)

type Slasher struct {
	MSPClient *msp.Client
	SDK       *fabsdk.FabricSDK
	Conf      *slasherConfig.Config
}

func New(conf *slasherConfig.Config) (*Slasher, error) {

	fabConfig := config.FromFile(conf.ConnectionProfile)
	sdk, err := fabsdk.New(fabConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new SDK: %s\n", err))
	}
	defer sdk.Close()

	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create msp client: %s\n", err))
	}

	return &Slasher{mspClient, sdk, conf}, nil
}

func (s *Slasher) EnrollUser(user, secret string) error {

	_, err := s.MSPClient.GetSigningIdentity(user)
	if err == msp.ErrUserNotFound {
		log.Info("Going to enroll user")
		err = s.MSPClient.Enroll(user, msp.WithSecret(secret))

		if err != nil {
			return errors.New(fmt.Sprintf("Failed to enroll user: %s\n", err))
		} else {
			log.Info("Success enroll user: %s\n", user)
			return nil
		}

	} else if err != nil {
		return errors.New(fmt.Sprintf("Failed to get user: %s\n", err))
	}

	log.Info(fmt.Sprintf("User %s already enrolled, skip enrollment.\n", user))
	return nil

}

func (s *Slasher) Revoke() (*msp.RevocationResponse, error) {

	RevocationResponse, err := s.MSPClient.Revoke(&msp.RevocationRequest{
		Name:   s.Conf.RevocationRequest.Name,
		Serial: s.Conf.RevocationRequest.Serial,
		AKI:    s.Conf.RevocationRequest.AKI,
		Reason: s.Conf.RevocationRequest.Reason,
		CAName: s.Conf.RevocationRequest.CAName,
	})
	if err != nil {
		return nil, err
	}

	return RevocationResponse, nil
}
