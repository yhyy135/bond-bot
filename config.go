package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token  string  `json:"token"`
	ChatId []int64 `json:"chat_id"`
}

func (config *Config) ReadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}
