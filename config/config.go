package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		DB DB `json:"db"`
	}

	DB struct {
		Address string `json:"address"`
		Port    string `json:"port"`
		User    string `json:"user"`
		Pass    string `json:"pass"`
	}
)

func Parse(cfg *Config, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(cfg)
}
