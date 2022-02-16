package config

import (
	"encoding/json"
	"os"
)

type ListenerConfig struct {
	Addr string `json:"addr"`
}

type MongoCfg struct {
	URL string `json:"url"`
}

type Config struct {
	MongoCfg       MongoCfg       `json:"mongo_cfg"`
	ListenerConfig ListenerConfig `json:"listener_config"`
}

func SetupConfig(filePath string, cfg *Config) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return err
	}

	return nil
}
