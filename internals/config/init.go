package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SourceDirectory string `json:"source_directory"`
	Entity          string `json:"entity"`
	Dto             string `json:"dto"`
	JsonController  string `json:"json_controller"`
	HtmlController  string `json:"html_controller"`
	Service         string `json:"service"`
	Repository      string `json:"repository"`
	Mapper          string `json:"mapper"`
	Transient       bool   `json:"-"`
}

var CONFIG_FILE_URL = os.Getenv("GOPATH") + "/bin/go-cap/resources/go-cap.json"
var RESOURCE_FOLDER = os.Getenv("GOPATH") + "/bin/go-cap/resources/"
var CONFIG *Config

func init() {
	_, err := os.Stat("./symfony.lock")

	if err != nil {
		panic("Command exectuded outside of a symfony root directory. Exiting")
	}

	file, err := os.ReadFile(CONFIG_FILE_URL)
	CONFIG = getConfig(file, err)
}

func getConfig(file []byte, err error) *Config {
	var config Config

	if err != nil {
		return handleConfigError()
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Couldn't parse config file. Switching to default values")
		return getDefaultConfig()
	}

	return &config
}

func handleConfigError() *Config {
	//mut
	var response string
	//mut
	config := getDefaultConfig()

	for response != "y" && response != "Y" && response != "n" && response != "N" {
		fmt.Println("Config file not found. Do you wish to create it ? (Y/N)")
		//mutation
		fmt.Scanln(&response)
	}

	if response == "y" || response == "Y" {
		err := os.MkdirAll(RESOURCE_FOLDER, 0777)

		if err != nil {
			panic("Couldn't create resource folder. Exiting")
		}

		newFile, err := os.Create(CONFIG_FILE_URL)
		defer newFile.Close()

		if err != nil {
			panic("Couldn't create config file. Exiting")
		}

		config = getDefaultConfig()
		confBytes, _ := json.MarshalIndent(config, "", "  ")
		newFile.Write(confBytes)
		//mutation
		config.Transient = false
	}

	return config
}

func getDefaultConfig() *Config {
	return &Config{
		SourceDirectory: "./src/",
		Entity:          "Entity/",
		Dto:             "Types/Dtos/",
		JsonController:  "Controller/Json/",
		HtmlController:  "Controller/Html/",
		Service:         "Service/",
		Transient:       true,
	}
}

func WriteConfigFile() {
	confBytes, _ := json.MarshalIndent(CONFIG, "", "  ")
	os.WriteFile(CONFIG_FILE_URL, confBytes, 0666)
}
