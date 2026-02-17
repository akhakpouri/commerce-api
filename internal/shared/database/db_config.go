package database

type DbConfig struct {
	Host     string
	User     string
	DbName   string
	Port     int
	Password string
	SSLMode  string
	Schema   string
}
