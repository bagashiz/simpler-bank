package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config contains and configures the environment variables for the application.
type Config struct {
	DB_HOST             string `mapstructure:"DB_HOST"`
	DB_USER             string `mapstructure:"DB_USER"`
	DB_PASSWORD         string `mapstructure:"DB_PASSWORD"`
	DB_NAME             string `mapstructure:"DB_NAME"`
	DB_PORT             string `mapstructure:"DB_PORT"`
	DB_SSL_MODE         string `mapstructure:"DB_SSL_MODE"`
	DB_TIMEZONE         string `mapstructure:"DB_TIMEZONE"`
	HTTP_SERVER_ADDRESS string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// NewConfig returns a new Config object
func NewConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

// GetDSN returns data source name for database in GORM format.
func (conf *Config) GetDSN() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.DB_HOST,
		conf.DB_USER,
		conf.DB_PASSWORD,
		conf.DB_NAME,
		conf.DB_PORT,
		conf.DB_SSL_MODE,
		conf.DB_TIMEZONE,
	)

	return dsn
}
