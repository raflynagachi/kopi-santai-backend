package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var Testing = "testing"

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type AppConfig struct {
	AppName                string
	BaseURL                string
	Port                   string
	ENV                    string
	JWTSecret              []byte
	JWTExpiredInMinuteTime int64
	DBConfig               dbConfig
}

var Config = AppConfig{
	AppName:                getEnv("APP_NAME", "Kopi Santai"),
	BaseURL:                getEnv("BASE_URL", "localhost"),
	Port:                   getEnv("PORT", "8080"),
	ENV:                    getEnv("ENV", Testing),
	JWTSecret:              []byte(getEnv("JWT_SECRET", "p@ssW0rd")),
	JWTExpiredInMinuteTime: 15,
	DBConfig: dbConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "kopi_santai_db"),
		Port:     getEnv("DB_PORT", "5432"),
	},
}

func getEnv(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}
