package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() error {
	connString := GetConnectionString()
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}

func GetConnectionString() string {
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"sqlserver://%s?database=%s&trusted_connection=true",
		host,
		database,
	)
}

// Helper to get env with default
func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
