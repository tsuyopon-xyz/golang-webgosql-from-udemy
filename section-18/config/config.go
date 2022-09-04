package config

import (
	"log"
	"myapp/utils"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

func init() {
	Init()
	utils.LoggingSetting(Config.LogFile)
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = ConfigList{
		Port:      os.Getenv("PORT"),
		SQLDriver: os.Getenv("SQL_DRIVER"),
		DbName:    os.Getenv("DB_NAME"),
		LogFile:   os.Getenv("LOG_FILE"),
	}
}
