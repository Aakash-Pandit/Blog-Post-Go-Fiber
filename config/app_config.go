package config

import (
	"os"

	"github.com/joho/godotenv"
)

func ReadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func GetEnv(env_variable string) string {
	if value, exists := os.LookupEnv(env_variable); exists {
		return value
	}

	panic("Key: " + env_variable + " variable does not exists.")
}

func GetApplicationConfig() *ApplicationConfig {
	ReadEnvFile()

	var (
		backend_port      = GetEnv("BACKEND_PORT")
		postgres_host     = GetEnv("POSTGRES_HOST")
		postgres_port     = GetEnv("POSTGRES_PORT")
		postgres_user     = GetEnv("POSTGRES_USER")
		postgres_password = GetEnv("POSTGRES_PASSWORD")
		postgres_db_name  = GetEnv("POSTGRES_DB_NAME")
	)

	configuration := &ApplicationConfig{
		BackendPort:      backend_port,
		PostgresHost:     postgres_host,
		PostgresPort:     postgres_port,
		PostgresUser:     postgres_user,
		PostgresPassword: postgres_password,
		PostgresDbName:   postgres_db_name,
	}

	return configuration
}
