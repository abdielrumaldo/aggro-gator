package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUser = user
	return write(*cfg)
}

func write(cfg Config) error {

	jsonData, err := json.Marshal(cfg) // Indent with 2 spaces
	if err != nil {
		fmt.Println("Error writting to configuration file: ", err)
		return err
	}
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error Getting Home Directory Path: ", err)
		return err
	}
	err = os.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error Writting to the configuration file", err)
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := homeDir + "/" + configFileName
	return configPath, nil
}

func Read() (Config, error) {

	config := Config{}
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error Getting Home Directory Path: ", err)
		return config, err
	}
	content, err := os.ReadFile(configFilePath)
	if err != nil {

		fmt.Printf("Error Reading file %s. Error:\n%s\n", configFilePath, err)
		return config, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Printf("Failed to parse JSON content in %s. Error:\n%s\n", configFilePath, err)
		return config, err
	}

	return config, nil
}
