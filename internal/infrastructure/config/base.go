package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// SetGlobalEnv set global env at the beginning
func SetGlobalEnv() {

	if viper.GetString("APP_ENV") == "develop" {
		gin.SetMode(gin.DebugMode)
	} else if viper.GetString("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

}

// SetConfig sets the configuration name and reads the config file.
func SetConfig(envPath string, envType string, configName string) error {
	viper.AddConfigPath(envPath)
	viper.SetConfigType(envType)
	viper.SetConfigName(configName)
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	return nil
}

// InitConfig initialing configs
func InitConfig(envPath string, envType string, configName string) {
	err := SetConfig(envPath, envType, configName)
	viper.SetDefault("PORT", 8080)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
