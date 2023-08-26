package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Server       serverConfiguration                 `yaml:"server"`
	Integrations map[string]integrationConfiguration `yaml:"integrations"`
}

type serverConfiguration struct {
	Port string `yaml:"port"`
}

type integrationConfiguration struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

func Load(file string) (*Configuration, error) {
	config := Configuration{}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
