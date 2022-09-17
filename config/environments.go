package config

import (
	"errors"
	"os"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/core"
)

func GetEnv(key string) (string, error) {

	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	return "", errors.New("Error: " + key + " env cannot be nil.")
}

func SetupEnv() *core.AppConfig {

	var (
		host, _        = GetEnv("DB_HOST")
		port, _        = GetEnv("DB_PORT")
		user, _        = GetEnv("DB_USER")
		dbName, _      = GetEnv("DB_NAME")
		password, _    = GetEnv("DB_PASSWORD")
		sslmode, _     = GetEnv("DB_SSLMODE")
		backendPort, _ = GetEnv("BACKEND_PORT")
	)

	return &core.AppConfig{
		DBHost:      host,
		DBPort:      port,
		DBUser:      user,
		DBName:      dbName,
		DBPassword:  password,
		DBSSLMode:   sslmode,
		BackendPort: backendPort,
	}
}
