package utils

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"dbname"`
}

func SaveConfig(config DBConfig, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

// Add any helper functions here
