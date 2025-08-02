package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var fileConfigName = ".gatorconfig.json"

type Config struct {
	DBURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func Read() Config {
	filepath := getFilepath()
	data, _ := os.ReadFile(filepath)
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUser = name
	filepath := getFilepath()
	data, _ := json.Marshal(cfg)
	return os.WriteFile(filepath, data, 0644)
}

func getFilepath() string {
	home, _ := os.UserHomeDir()
	filepath := filepath.Join(home, fileConfigName)
	return filepath
}
