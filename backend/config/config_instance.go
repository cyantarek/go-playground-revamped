package config

var Cfg = Config{
	HttpPort:         EnvOrDefault("HTTP_PORT", "5000"),
	APIPort:          EnvOrDefault("API_PORT", "5001"),
	GRPCPort:         EnvOrDefault("GRPC_PORT", "5002"),
	DBHost:           EnvOrDefault("DB_PORT", "localhost"),
	DBPort:           EnvOrDefault("DB_PORT", "9432"),
	DBName:           EnvOrDefault("DB_NAME", "boilerplate"),
	DBUser:           EnvOrDefault("DB_USER", "listmonk"),
	DBPassword:       EnvOrDefault("DB_PASSWORD", "listmonk"),
	JWTSecret:        EnvOrDefault("JWT_SECRET", "*234230KJHDS"),
	JWTRefreshSecret: EnvOrDefault("JWT_REFRESH_SECRET", "WE13123??__<>"),
	AuthSkipper: map[string]bool{
		"/playground.Playground/Ping":       true,
		"/playground.Playground/FormatCode": true,
		"/playground.Playground/RunCode":    true,
		"/playground.Playground/ShareCode":  true,
	},
	RedisHost:   "localhost",
	RedisPort:   "6379",
	GRPCWebHost: "0.0.0.0",
	GRPCWebPort: EnvOrDefault("GRPC_WEB_PORT", "5003"),
}
