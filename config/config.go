package config

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type DBConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
}
type APIConfig struct {
	Port      int `yaml:"port"`
	ServerMode string `yaml:"ginServerMode"`
}
type Config struct {
	APIConfig APIConfig `yaml:"apiConfig"`
	DBConfig  DBConfig  `yaml:"db"`
}

var c Config

type envFn func(string) error

var envFuncMapper = map[string]envFn{
	"DATABASE_HOST":     overrideDatabaseHost,
	"DATABASE_USER":     overrideDatabaseUser,
	"DATABASE_PORT":     overrideDatabasePort,
	"DATABASE_NAME":     overrideDatabaseName,
	"DATABASE_PASSWORD": overrideDatabasePassword,
	"PORT": overrideAPIPort,
}

func New() *Config {
	env := os.Getenv("APP_ENV")
	var envCfg Config

	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		logrus.Errorf("Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		logrus.Errorf("Unmarshal: %v", err)
	}
	if env != "" {
		yamlFile, err = ioutil.ReadFile(fmt.Sprintf("./config/config.%s.yaml", env))
		if err != nil {
			logrus.Warnf("config.%s.yaml .Get err %v ", env, err)
		}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			logrus.Errorf("Unmarshal: %v", err)
		}

		if err = mergo.MergeWithOverwrite(&c, envCfg); err != nil {
			logrus.Errorf("Merge configs: %v", err)
		}
	}
	//override with env variables
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if op, ok := envFuncMapper[pair[0]]; ok {
			err = op(pair[1])
			if err != nil {
				panic(err)
			}
		}
	}
	return &c
}

// Override database host env if variable exists
func overrideDatabaseHost(host string) error {
	if host != "" {
		c.DBConfig.Host = host
		return nil
	}
	return fmt.Errorf("DATABASE HOST ENV NOT FOUND")
}

// Override database host env if variable exists
func overrideDatabaseUser(user string) error {
	if user != "" {
		c.DBConfig.User = user
		return nil
	}
	return fmt.Errorf("DATABASE USER ENV NOT FOUND")
}

// Override database port env if variable exists
func overrideDatabasePort(port string) error {
	if port != "" {
		portInt, err := strconv.ParseUint(port, 10, 32)
		if err != nil {
			return err
		}
		c.DBConfig.Port = uint(portInt)
		return nil
	}
	return fmt.Errorf("DATABASE PORT ENV NOT FOUND")
}

// Override database name env if variable exists
func overrideDatabaseName(name string) error {
	if name != "" {
		c.DBConfig.Name = name
		return nil
	}
	return fmt.Errorf("DATABASE NAME ENV NOT FOUND")
}

// Override database password env if variable exists
func overrideDatabasePassword(password string) error {
	if password != "" {
		c.DBConfig.Password = password
		return nil
	}
	return fmt.Errorf("DATABASE PASSWORD ENV NOT FOUND")
}
// Override database password env if variable exists
func overrideAPIPort(port string) error {
	if port != "" {
		portInt, err := strconv.ParseInt(port, 10, 32)
		if err != nil {
			return err
		}
		c.APIConfig.Port = int(portInt)
		return nil
	}
	return fmt.Errorf("API PORT ENV NOT FOUND")
}
