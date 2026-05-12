package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Listen    string `json:"listen"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IPAddress  string `json:"ip_address"`
	PublicPort string `json:"public_port"`
	Smtp       string `json:"smtp"`
	StoreEvery int    `json:"store_every"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
