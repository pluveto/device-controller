package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	API     string `json:"api"`
	AuthKey string `json:"authKey"`
}

func loadConfig(fileName string) (*Config, error) {
	var conf Config
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
