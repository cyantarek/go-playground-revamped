package config

// Config holds the configuration info in a place to be accessed later during runtime
type Config struct {
	Server struct {
		Grpc struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"grpc"`
		GrpcWeb struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"grpc_web"`
	} `yaml:"server"`
	Db struct {
		Mongo struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DbName   string `yaml:"db_name"`
		} `yaml:"mongo"`
	} `yaml:"db"`
	SMTP struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
	Security struct {
		TLS struct {
			KeyFile  string `yaml:"key_file"`
			CertFile string `yaml:"cert_file"`
		} `yaml:"tls"`
	} `yaml:"security"`
	Worker struct {
		Count int `yaml:"count"`
	} `yaml:"worker"`
}
