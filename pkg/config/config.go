package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// InitConfig init viper and return it.
func InitConfig(filepath string) *viper.Viper {
	viper.AddConfigPath(filepath)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return viper.GetViper()
		}

		log.Fatalln("failed to read config file:", err)
	}
	return viper.GetViper()
}

// SetupDefault setup default value and return error if required fields is not present.
func SetupDefault(v *viper.Viper) error {
	if !v.IsSet("token") {
		return fmt.Errorf("token is required")
	}
	if !v.IsSet("repo") {
		return fmt.Errorf("repo is required")
	}
	v.SetDefault("arch", "amd64")
	v.SetDefault("tag", "latest")

	return nil
}
