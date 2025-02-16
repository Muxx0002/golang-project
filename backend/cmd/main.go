package main

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/transport/routes"
	"github.com/Muxx0002/golang-project/tree/main/backend/pkg/tools"
	"github.com/google/logger"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config/config.env")
	viper.ReadInConfig()
	file := tools.CreateLogFile()
	defer logger.Init("logger", false, true, file).Close()
	postgres.InitDB()
	routes.Routes()
}
