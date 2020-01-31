package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
	Database struct {
		Port     string `json:"port"`
		Host     string `json:"host"`
		Username string `json:"username"`
		Dbname   string `json:"dbname"`
		Password string `json:"password"`
	} `json:"database"`
}

// Initializes the configurations for the server
func InitConfig(isDev bool) Config {
	var config Config
	var configFile *os.File
	var err error

	// Decides if the environment is dev or not
	if isDev {
		configFile, err = os.Open("config/dev.config.json")
	} else {
		configFile, err = os.Open("config/config.json")
	}

	// Defaults to closing if it does not work
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Parses the config.json
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
