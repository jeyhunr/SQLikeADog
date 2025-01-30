package auth

import (
	"encoding/json"
	"os"
)

type Credentials struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

// SaveCredentials saves credentials to a JSON file
func SaveCredentials(creds Credentials) error {
	file, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("internal/auth/credentials.json", file, 0644)
}

// LoadCredentials loads credentials from a JSON file
func LoadCredentials() (Credentials, error) {
	var creds Credentials
	file, err := os.ReadFile("internal/auth/credentials.json")
	if err != nil {
		return creds, err
	}
	err = json.Unmarshal(file, &creds)
	return creds, err
}

// DeleteCredentials removes the credentials JSON file
func DeleteCredentials() error {
	return os.Remove("internal/auth/credentials.json")
}
