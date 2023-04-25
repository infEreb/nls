package api

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ApiConfig struct {
	Api struct {
		Port     int               `yaml:"port"`
		LogFile  string            `yaml:"logFile"`
		Gateways map[string]string `yaml:"gateways,omitempty"`
		Type string `yaml:"type,omitempty"`
	} `yaml:"api"`
}

func ParseConfig(yamlConf string) (*ApiConfig, error) {
	conf := &ApiConfig{}
	confBytes, err := os.ReadFile(yamlConf)
	if err != nil {
		return nil, fmt.Errorf("error to read config file: %s", err)
	}

	if err := yaml.NewDecoder(bytes.NewReader(confBytes)).Decode(conf); err != nil {
		return nil, fmt.Errorf("error to decode config file: %s", err)
	}

	return conf, nil
}