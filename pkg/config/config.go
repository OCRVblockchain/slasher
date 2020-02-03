package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	ActionUser        string
	Secret            string
	ConnectionProfile string
	RevocationRequest
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

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// parse flags
	confpath := viper.Get("configpath")
	identity := viper.Get("identity")
	secret := viper.Get("secret")

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

	slasherConfiguration.ActionUser = identity.(string)
	slasherConfiguration.Secret = secret.(string)

	return slasherConfiguration, nil
}
