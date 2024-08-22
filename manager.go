package oficonnectbot

import (
	"encoding/json"
	"log"
	"os"
)

type BotConfig struct {
	Excludes []string `json:"excludes"`
}

const FILEPATH = "/home/csoria/.oficonnectbot.json"

func ReadConfig() *BotConfig {
	content, err := os.ReadFile(FILEPATH)

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

	os.WriteFile(FILEPATH, data, 0644)
}
