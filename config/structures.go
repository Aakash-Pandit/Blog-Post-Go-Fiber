package config

type ApplicationConfig struct {
	BackendPort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDbName   string
}
