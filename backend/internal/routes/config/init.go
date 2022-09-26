package config

import (
	"github.com/spf13/viper"
)

// InitConfig gets full config from file
func InitConfig() (*FileConfig, error) {
	var newConfig *FileConfig

	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	err := viper.ReadInConfig()
	if err != nil {
		return newConfig, err
	}

	newConfig = &FileConfig{
		Port: viper.GetString("port"),
	}

	return newConfig, nil
}
