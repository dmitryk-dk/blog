package config

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	DbName       string `json:"dbName"`
	DbDriverName string `json:"dbDriver"`
}

func Read(fileName string) (*Config, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
