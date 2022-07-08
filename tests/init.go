package tests

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var configFile = flag.String("f", "test_config.yaml", "the config file")

var testConfig = ReadConfig(configFile)

type Config struct {
	Database struct {
		DSN   string `yaml:"dsn"`
		Redis string `yaml:"redis"`
	}
	Email struct {
		Account  string `yaml:"account"`
		Password string `yaml:"password"`
		To       string `yaml:"to"`
	}
}

func ReadConfig(path *string) *Config {
	yamlFile, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil
	}
	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil
	}
	return config
}
