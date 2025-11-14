package config

import (
	"fmt"
	"os"
)

type Config struct {
	DbURL string `json:"db_url"`
}

func (c *Config) Read() error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	fmt.Printf("%s is the home directory", homeDir)
	return nil
}
