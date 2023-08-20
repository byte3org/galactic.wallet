package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Environment string `json:"env"`
	Port        int    `json:"port"`
	DbString    string `json:"db_string"`
	DbName      string `json:"db_name"`
	AuthService string `json:"auth_service"`
	PublicKey   string `json:"public_key"`
}

var config *Config

func intitialize() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("[x] error : couldnt read config file", err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("[x] error", err.Error())
	}
	config = &conf
}

func GetConfig() *Config {
	return config
}
