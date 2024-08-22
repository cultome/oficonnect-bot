package oficonnectbot

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type BotConfig struct {
	Excludes []string `json:"excludes"`
}

func ReadConfig() *BotConfig {
	content, err := os.ReadFile(ConfigFilePath())

	if err != nil {
		log.Fatalf("Unable to read config file: %s", err.Error())
	}

	var config BotConfig
	json.Unmarshal(content, &config)

	return &config
}

func (c *BotConfig) Persist() {
	data, err := json.Marshal(c)

	if err != nil {
		log.Fatalf("Unable to read config file")
	}

	os.WriteFile(ConfigFilePath(), data, 0644)
}

func ConfigFilePath() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatalf("Unable to get user directory")
	}

	return filepath.Join(usr.HomeDir, ".oficonnectbot.json")
}
