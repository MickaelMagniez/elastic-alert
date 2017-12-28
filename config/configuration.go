package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigurationElastic struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ConfigurationTargetsEmailSMTP struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ConfigurationTargetsEmail struct {
	Sender string                        `json:"sender"`
	Smtp   ConfigurationTargetsEmailSMTP `json:"smtp"`
}
type ConfigurationTargets struct {
	Email ConfigurationTargetsEmail `json:"email"`
}
type Configuration struct {
	Targets ConfigurationTargets `json:"targets"`
	Elastic ConfigurationElastic `json:"elastic"`
}

var configuration *Configuration

func GetConfiguration() Configuration {
	return *configuration
}

func InitConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&configuration)
}
