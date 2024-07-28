package config

type Config struct {
	ApiVersion string
	SrvTimeout int
	SrvPort int
}

func NewConfig() *Config {
	return &Config{
		ApiVersion: "v1",
		SrvTimeout: 15,
		SrvPort: 8080,
	}
}