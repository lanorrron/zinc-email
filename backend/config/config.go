package config

import "os"

type Config struct {
	ZincSearchHost      string
	ZincSearchIndexName string
	ZincSearchUser      string
	ZincSearchPassword  string
	EmailDirectory      string
	ModeIndex           string
	ServerPort          string
}

func LoadConfig() *Config {
	return &Config{
		ZincSearchHost:      os.Getenv("ZINC_SEARCH_HOST"),
		ZincSearchIndexName: os.Getenv("ZINC_SEARCH_INDEX_NAME"),
		ZincSearchUser:      os.Getenv("ZINC_SEARCH_USER"),
		ZincSearchPassword:  os.Getenv("ZINC_SEARCH_PASSWORD"),
		EmailDirectory:      os.Getenv("EMAIL_DIRECTORY"),
		ModeIndex:           os.Getenv("MODE_INDEX"),
		ServerPort:          os.Getenv("SERVER_PORT"),
	}
}
