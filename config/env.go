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
	DBPass     string
	DBAddress  string
	DBName     string
}
 
var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()
	return Config{ 
		PublicHost: getEnv("PUBLIC_HOST", "https://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", ""),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "restapi"), 
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}