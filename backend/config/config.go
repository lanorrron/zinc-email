package config

import "os"

type Config struct {
	ZincSearchHost     string
	ZincSearchUser     string
	ZincSearchPassword string
	EmailDirectory     string
	ServerPort         string
}

func LoadConfig() *Config {
	return &Config{
		ZincSearchHost:     os.Getenv("ZINC_SEARCH_HOST"),
		ZincSearchUser:     os.Getenv("ZINC_SEARCH_USER"),
		ZincSearchPassword: os.Getenv("ZINC_SEARCH_PASSWORD"),
		EmailDirectory:     os.Getenv("EMAIL_DIRECTORY"),
		ServerPort:         os.Getenv("SERVER_PORT"),
	}
}
