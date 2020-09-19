package config

import "os"

// Configuration object for app wide config
type Configuration struct {
	AppName string
	Server  Server
	Logger  Logger
}

// Server config
type Server struct {
	Host string
	Port string
}

// Logger config
type Logger struct {
	Level  string
	Format string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// LoadConfig loads configuration file
func LoadConfig() *Configuration {

	config := &Configuration{
		AppName: getEnv("NAME", "seed-go-template"),
		Server: Server{
			Host: getEnv("HOST", "0.0.0.0"),
			Port: getEnv("PORT", "7000"),
		},
		Logger: Logger{
			Level:  "info",
			Format: "",
		},
	}

	return config
}
