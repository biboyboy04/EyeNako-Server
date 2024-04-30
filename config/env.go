package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTExpirationInSeconds int64
	JWTSecret string
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
		JWTSecret: getEnv("JWT_SECRET", "sikretong-malupet-pwedeng-pabulong?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600 * 24 * 7 ),
	}
}

func getEnv(key, fallback string) string  {
	if value, ok := os.LookupEnv(key);  ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key);  ok {
		// Parse just to make sure value is int
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}