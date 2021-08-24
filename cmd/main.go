package main

import (
    "github.com/spf13/viper"
)

func initConfig() error {
    viper.AddConfigPath("config")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}

func main() {
    if err := initConfig(); err != nil {
        // log into database
		
    }

}

