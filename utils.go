package main

import (
	"encoding/json"
	"os"
	"sample_golang_application/common"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func init() {
	common.Info.Println("initialising")
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		common.Fatal.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		common.Fatal.Fatalln("Cannot get configuration from file", err)
	}
}
