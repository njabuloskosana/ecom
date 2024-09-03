package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DbUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default or environment variables")
	}
	log.Println(".env file found")
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", ""),
		Port:       getEnv("PORT", ""),
		DbUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress:  getEnv("DB_HOST", ""), // PostgreSQL default port
		DBName:     getEnv("DB_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
