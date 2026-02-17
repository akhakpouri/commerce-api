package managers

import (
	"commerce/internal/shared/database"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

func NewDbConfig(filePath string) (database.DbConfig, error) {
	slog.Info("Loading database configuration from file:", "filePath", filePath)
	// Load database configuration from the specified file if it exists
	cfg, err := dbConfigFromFile(filePath)
	if err != nil {
		slog.Error("Error loading config:", "error", err)
		//otherwise, load from environment variables if they exist
		cfg = getConfigFromEnv()
		return cfg, err
	}

	return cfg, nil
}

func getConfigFromEnv() database.DbConfig {
	cfg := database.DbConfig{}
	if host, ok := os.LookupEnv("DB_HOST"); ok {
		cfg.Host = host
	}

	if user, ok := os.LookupEnv("DB_USER"); ok {
		cfg.User = user
	}

	if dbName, ok := os.LookupEnv("DB_NAME"); ok {
		cfg.DbName = dbName
	}

	if password, ok := os.LookupEnv("DB_PASSWORD"); ok {
		cfg.Password = password
	}

	if sslMode, ok := os.LookupEnv("DB_SSLMODE"); ok {
		cfg.SSLMode = sslMode
	}

	if schema, ok := os.LookupEnv("DB_SCHEMA"); ok {
		cfg.Schema = schema
	}

	cfg.Port = getPortFromEnv()
	return cfg
}

func getPortFromEnv() int {
	const defaultPort = 5432
	if port, ok := os.LookupEnv("DB_PORT"); ok {
		p, err := fmt.Sscanf(port, "%d", &port)
		if err != nil {
			slog.Error("Invalid port value in environment variable:", "error", err)
			return defaultPort // default port
		}
		return p
	}
	return defaultPort // default port
}

func dbConfigFromFile(filePath string) (database.DbConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("failed to open config file:", "error", err)
		return database.DbConfig{}, fmt.Errorf("failed to open config file: %w", err)
	}
	//reading only - safe to ignore.
	defer func() { _ = file.Close() }()

	var cfg database.DbConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		slog.Error("failed to decode config file:", "error", err)
		return database.DbConfig{}, fmt.Errorf("failed to decode config file: %w", err)
	}

	return cfg, nil
}
