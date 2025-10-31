package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Database string
	Port     int
}

func getEnv(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	log.Printf("AVISO: variável %s não está definida, usando valor padrão: %s", key, defaultVal)
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("Erro ao converter %s=%s para int. Usando padrão: %d", key, valStr, defaultVal)
		return defaultVal
	}
	return val
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente")
	}

	database := getEnv("DATABASE", "")
	if strings.TrimSpace(database) == "" {
		return nil, fmt.Errorf("variável DATABASE não definida")
	}

	port := getEnvAsInt("API_PORT", 9000)

	env := getEnv(strings.TrimSpace(strings.ToUpper("ENV")), "DEVELOPMENT")

	return &Config{
		Env:      env,
		Database: database,
		Port:     port,
	}, nil
}
