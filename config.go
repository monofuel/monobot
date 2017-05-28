package monobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//SettingOptions is the struct that controls what the config file looks like
type SettingOptions struct {
	IrcServers       []string `json:"ircServers"`
	Owner            string   `json:"owner"`
	IrcPrefix        string   `json:"ircPrefix"`
	DiscordToken     string   `json:"discordToken"`
	PushbulletAPIKey string   `json:"pushbulletAPIKey"`
}

//ConfigFilename is the name of the configuration file
var ConfigFilename = "config.json"

//Settings is the object detailing the loaded configuration
var Settings = &SettingOptions{
	IrcServers:   []string{"japura.net:6667"},
	Owner:        "monofuel",
	IrcPrefix:    "$mono",
	DiscordToken: "",
}

func Configuration(s *SettingOptions) {
	Settings = s
}

func loadConfig() error {

	configContents, err := ioutil.ReadFile(ConfigFilename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("generating new config")

			_, err = os.Create(ConfigFilename)
			if err != nil {
				log.Fatalf("error creating config: %s\n", err.Error())
			}

			bytes, err := json.Marshal(&Settings)
			if err != nil {
				fmt.Printf("Error marshalling new defaults: %s\n", err.Error())
			}
			err = ioutil.WriteFile(ConfigFilename, bytes, 0755)
			if err != nil {
				fmt.Printf("Error writing defaults: %s\n", err.Error())
			}
		} else {
			log.Fatalf("error loading config: %s\n", err.Error())
		}
	} else {

		if err = json.Unmarshal(configContents, &Settings); err != nil {
			fmt.Printf("Error parsing file: %s\n", err.Error())
		}
	}
	/*
		fmt.Println("configuration:", configuration)
		fmt.Println("loaded existing configuration")
		if err = setUnassignedValues(Settings, configuration); err != nil {
			return err
		}
		bytes, err := json.Marshal(&Settings)
		if err != nil {
			fmt.Printf("Error marshalling new defaults: %s\n", err.Error())
		}
		/*
		err = ioutil.WriteFile(ConfigFilename, bytes, 0755)
		if err != nil {
			fmt.Printf("Error writing new defaults: %s\n", err.Error())
		}
		fmt.Println("saving configuration with new values")
	*/
	return nil
}

/*
func setUnassignedValues(dst interface{}, src interface{}) error {
	s1 := reflect.ValueOf(dst)
	if s1.Kind() != reflect.Struct {
		return fmt.Errorf("destination param was not a struct")
	}

	s2 := reflect.ValueOf(src)
	if s2.Kind() != reflect.Struct {
		return fmt.Errorf("source param was not a struct")
	}

	for i := 0; i < s2.NumField(); i++ {
		f := s2.Field(i)
		if f.Kind() != reflect.String {
			if f.IsNil() {
				fmt.Println("skipping nil value")
				continue
			}
		}
		if f.Kind() == reflect.Slice {
			fmt.Println("skipping slice")
		} else {
			fmt.Println("setting ", f.String())
			s1.Set(f)
		}
	}
	return nil
}*/
