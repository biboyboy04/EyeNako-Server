package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config{

	// This is needed for setting the env values as without this, 
	// the .env values are not being read
	godotenv.Load()

	return Config {
		PublicHost: getEnv("PUBLIC_HOST", "http:/localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"), 
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), 
		getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "go-test"),
	}
}

func getEnv(key, fallback string) string  {
	if value, ok := os.LookupEnv(key);  ok {
		return value
	}
	return fallback
}