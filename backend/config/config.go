package config

import (
	"os"
)

type Config struct {
	HttpPort          string
	GRPCPort          string
	APIPort           string
	TLSCertFile       string
	TLSKeyFile        string
	DBHost            string
	DBPort            string
	DBName            string
	DBUser            string
	DBPassword        string
	JWTSecret         string
	JWTRefreshSecret  string
	MailgunDomain     string
	MailgunAPIKey     string
	AuthSkipper       map[string]bool
	GCPServiceAccount string
	RedisHost         string
	RedisPort         string
	GRPCWebHost       string
	GRPCWebPort       string
}

func EnvOrDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}
