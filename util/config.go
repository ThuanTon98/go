package util

import "github.com/spf13/viper"

// Config store all configuration of the application
// The values are read by viper from a config file or environment variables.

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSOURCE      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or enviroment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}