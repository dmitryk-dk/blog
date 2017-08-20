package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type Config struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	DbName       string `json:"dbName"`
	DbDriverName string `json:"dbDriver"`
}

var config *Config

func GetConfig () *Config {
	if config == nil {
		config = new(Config).readConfig()
	}
	return config
}

func (c Config) readConfig() (*Config) {
	flagName := c.readFlags()
	file, err := ioutil.ReadFile(*flagName)
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil
	}
	return config
}

func (c Config) readFlags () (configFile *string) {
	configFile = flag.String("dbconfig", "dbconfig.json", "Path to config file")
	flag.Parse()
	return
}
