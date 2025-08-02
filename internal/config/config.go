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
	path := getFilepath()
	data, _ := os.ReadFile(path)
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUser = name
	path := getFilepath()
	data, _ := json.Marshal(cfg)
	return os.WriteFile(path, data, 0644)
}

func getFilepath() string {
	home, _ := os.UserHomeDir()
	filepath := filepath.Join(home, fileConfigName)
	return filepath
}
