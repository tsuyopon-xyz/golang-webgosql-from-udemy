package config

import (
	"log"
	"myapp/utils"
	"os"

	"github.com/joho/godotenv"
)

type Db struct {
	DbPort     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
}

type ConfigList struct {
	Port    string
	LogFile string
	Db
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
		Port:    os.Getenv("PORT"),
		LogFile: os.Getenv("LOG_FILE"),
		Db: Db{
			DbPort:     os.Getenv("DB_PORT"),
			DbHost:     os.Getenv("DB_HOST"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
		},
	}
}
