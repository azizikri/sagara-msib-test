package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    App struct {
        Port      int
    }
    Database struct {
        Driver string
        URL    string
    }
}

var AppConfig Config

func InitConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }
}
