package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func jsonInit() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		defaultConfig := Config{ListenPort: "8443", CertFile: "webhook.crt", KeyFile: "webhook.key"}
		fmt.Println("Config file not found. Generating one...")
		err = setConfig(defaultConfig)
		if err != nil {
			fmt.Println("Couldn't create default config file: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Default config file generated. Please configure it and run again.")
		os.Exit(2)
	}

	if _, err := os.Stat("bots.json"); os.IsNotExist(err) {
		defaultBots := Bots{Bot: []Bot{Bot{URI: "/webhook", Host: "https://192.168.0.100:443"}}}
		fmt.Println("Bots file not found. Generating one...")
		err = setBots(defaultBots)
		if err != nil {
			fmt.Println("Couldn't create default bots file: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Default bots file generated. Please configure it and run again.")
		os.Exit(3)
	}
}

func getConfig() (config Config, err error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return config, errors.New("could not get config: " + err.Error())
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, errors.New("could not get config: " + err.Error())
	}
	return config, nil
}

func setConfig(config Config) error {
	marsh, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return errors.New("could not set config: " + err.Error())
	}
	err = ioutil.WriteFile("config.json", marsh, os.ModePerm)
	if err != nil {
		return errors.New("could not set config: " + err.Error())
	}
	return nil
}

func getBots() (bots Bots, err error) {
	file, err := ioutil.ReadFile("bots.json")
	if err != nil {
		return bots, errors.New("could not get bots: " + err.Error())
	}
	err = json.Unmarshal(file, &bots)
	if err != nil {
		return bots, errors.New("could not get bots: " + err.Error())
	}
	return bots, nil
}

func setBots(bots Bots) error {
	marsh, err := json.MarshalIndent(bots, "", "\t")
	if err != nil {
		return errors.New("could not set bots: " + err.Error())
	}
	err = ioutil.WriteFile("bots.json", marsh, os.ModePerm)
	if err != nil {
		return errors.New("could not set bots: " + err.Error())
	}
	return nil
}
