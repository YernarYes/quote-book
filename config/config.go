package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	HTTP     HTTP     `json:"http"`
	Postgres Postgres `json:"postgres"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Postgres struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}

func ConfigInit() (*Config, error) {
	file, err := os.Open("./config/config.json")
	fmt.Println()
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %v", err)
	}

	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config %v", err)
	}

	return &config, nil
}
