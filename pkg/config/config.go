package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Mode              string
	ActionUser        string
	Secret            string
	ConnectionProfile string
	RevocationRequest
	RemoveIdentityRequest
}

type RemoveIdentityRequest struct {
	ID     string
	Force  bool
	CAName string
}

type RevocationRequest struct {
	Name   string
	Serial string
	AKI    string
	Reason string
	CAName string
}

func GetConfig() (*Config, error) {

	var slasherConfiguration *Config

	flag.String("configpath", "./pkg/config/", "path to Slasher config folder")
	flag.String("identity", "admin", "identity that do revocation request")
	flag.String("secret", "", "identity secret")
	flag.String("mode", "fullslash", "mode\n--mode revokecert - for certificate revocation\n--mode removeidentity - for identity removal\n--mode fullslash - for both options")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// parse flags
	confpath := viper.Get("configpath")
	identity := viper.Get("identity")
	secret := viper.Get("secret")
	mode := viper.Get("mode")

	// read config
	viper.SetConfigName("config")
	viper.AddConfigPath(confpath.(string))
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New(fmt.Sprintf("unable to read config file, %s", err))
	}

	err := viper.Unmarshal(&slasherConfiguration)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to decode into struct, %v", err))
	}

	m, ok := mode.(string)
	if !ok {
		return nil, errors.New("Choose mode:\n--mode revokecert - for certificate revocation\n--mode removeidentity - for identity removal\n--mode fullslash - for both options")
	}

	i, ok := identity.(string)
	if !ok {
		return nil, errors.New("Choose identity:\n--identity myidentity\nIT'S YOUR IDENTITY, NOT FOR REMOVAL")
	}

	s, ok := secret.(string)
	if !ok {
		return nil, errors.New("Type secret for your identity:\n--secret mysecret\n")
	}

	slasherConfiguration.Mode = m
	slasherConfiguration.ActionUser = i
	slasherConfiguration.Secret = s

	return slasherConfiguration, nil
}
