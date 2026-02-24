package database

type DbConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	DbName   string `json:"dbname"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
	Schema   string `json:"schema"`
}
