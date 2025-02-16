package main

import "github.com/spf13/viper"

func main() {
	viper.SetConfigFile("config/config.env")
	viper.ReadInConfig()
}
