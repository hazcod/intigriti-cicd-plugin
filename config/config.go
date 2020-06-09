package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	IntigritiClientID		string		`yaml:"intigriti_client_id"`
	IntigritiClientSecret	string		`yaml:"intigriti_client_secret"`

	Tresholds struct {
		MaxCritical 		uint		`yaml:"critical"`
		MaxHigh 			uint 		`yaml:"high"`
		MaxMedium			uint		`yaml:"medium"`
		MaxLow				uint		`yaml:"low"`
	} `yaml:"tresholds"`
}

func ParseConfig(configPath string) (Config, error) {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, errors.Wrap(err, "could not read configuration file")
	}

	var config Config
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, errors.Wrap(err, "invalid yaml configuration")
	}

	return config, nil
}