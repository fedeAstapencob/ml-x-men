package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DBConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
}
type Config struct {
	DBConfig DBConfig `yaml:"db"`
}

var c Config

func New() *Config {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		logrus.Errorf("Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		logrus.Errorf("Unmarshal: %v", err)
	}
	return &c
}
