package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port      string `yaml:"port"`
	SecretKey string `yaml:"secretKey"`
	MongoUrl  string `yaml:"mongoUrl"`
}

func NewConfig(path string) (*Config, error) {
	c := new(Config)
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %v", err)
	}

	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		return nil, fmt.Errorf("cannot unmarshal config file: %v", err)
	}

	return c, nil
}
