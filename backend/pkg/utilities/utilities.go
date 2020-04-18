package utilities

import (
	"backend/config"
	"encoding/json"
	"os"
)

func GetEnv(key, defaults string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaults
}

func PrintConfig(cfg *config.Config) []byte {
	str, err := json.MarshalIndent(cfg, " ", " ")
	if err != nil {
		return nil
	}
	return str
}
