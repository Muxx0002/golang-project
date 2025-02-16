package main

import (
	"project/pkg/tools"

	"github.com/google/logger"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config/config.env")
	viper.ReadInConfig()
	file := tools.CreateLogFile()
	defer logger.Init("logger", false, true, file).Close()
}
