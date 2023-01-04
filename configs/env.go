package configs

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// GetDSN returns data source name for database connection with specified environment variables from .env file.
func GetDSN() (dsn string) {
	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	return
}

// GetHTTPServerAddress returns HTTP server address for routers connection with specified environment variables from .env file.
func GetHTTPServerAddress() (address string) {
	address = os.Getenv("HTTP_SERVER_ADDRESS")
	return
}
