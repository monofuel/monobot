package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type SettingOptions struct {
	IrcServers []string `json:"ircServers"`
	ApiServer  bool     `json:"apiServer"`
}

var CONFIG_FILENAME = "config.json"
var defaults = SettingOptions{
	IrcServers: []string{"japura.net"},
	ApiServer:  false,
}
var Settings SettingOptions

func loadConfig() {
	file, err := os.Open(CONFIG_FILENAME)
	if err != nil {
		if err == os.ErrNotExist {
			fmt.Println("generating new config")

			file, err = os.Create(CONFIG_FILENAME)
			if err != nil {
				log.Fatalf("error creating config: %s\n", err.Error())
			}
		} else {
			log.Fatalf("error loading config: %s\n", err.Error())
		}
	}
	if err := json.NewDecoder(file).Decode(&Settings); err != nil {
		log.Fatalf("Error parsing file: %s\n", err.Error())
	}
}
