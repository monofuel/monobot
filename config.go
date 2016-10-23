package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

//SettingOptions is the struct that controls what the config file looks like
type SettingOptions struct {
	IrcServers []string `json:"ircServers"`
	APIServer  bool     `json:"apiServer"`
	Owner      string   `json:"owner"`
}

//ConfigFilename is the name of the configuration file
var ConfigFilename = "config.json"

//Settings is the object detailing the loaded configuration
var Settings = SettingOptions{
	IrcServers: []string{"japura.net:6667"},
	APIServer:  false,
	Owner:      "monofuel",
}

func loadConfig() error {

	file, err := os.Open(ConfigFilename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("generating new config")

			file, err = os.Create(ConfigFilename)
			if err != nil {
				log.Fatalf("error creating config: %s\n", err.Error())
			}
			if err = json.NewEncoder(file).Encode(&Settings); err != nil {
				fmt.Printf("Error setting new defaults: %s\n", err.Error())
			}
		} else {
			log.Fatalf("error loading config: %s\n", err.Error())
		}
	}
	defer file.Close()
	configuration := make(map[string]interface{})
	if err = json.NewDecoder(file).Decode(&configuration); err != nil {
		fmt.Printf("Error parsing file: %s\n", err.Error())
	}
	fmt.Println("loaded existing configuration")
	if err = setUnassignedValues(Settings, configuration); err != nil {
		return err
	}
	file.Truncate(0)
	if err := json.NewEncoder(file).Encode(&Settings); err != nil {
		fmt.Printf("Error updating new defaults: %s\n", err.Error())
	}
	fmt.Println("saving configuration with new values")
	return nil
}

func setUnassignedValues(t interface{}, m map[string]interface{}) error {
	s := reflect.ValueOf(t)
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("first param was not a struct")
	}
	for key, value := range m {
		f := s.FieldByName(key)
		if !f.IsValid() {
			return fmt.Errorf("field missing from struct: %s\n", key)
		}
		if !f.CanSet() {
			return fmt.Errorf("could not set field on struct: %s\n", key)
		}
		f.Set(reflect.ValueOf(value))
	}
	return nil
}
