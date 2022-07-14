package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Set up Configuration type
type Config struct {
	Port string `mapstructure:"port"`
	Conn string `mapstructure:"connection_string"`
}

// AppConfig loads through the pointer of a Config struct
var AppConfig *Config

// LoadConfig loads the config from the given file
func LoadConfig() error {
	// Load the config from the file
	log.Info("Loading config...")

	// use Viper to load the config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// Try and read the config
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Marshall in the config file
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal("Error unmarshalling config file: ", err)
		return err
	}

	// If we're here, we're good
	return nil
}
