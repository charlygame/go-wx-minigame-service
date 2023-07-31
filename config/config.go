package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	AppName  string
	Host     string
	MongoURI string
}

var config appConfig

func InitWithEnvFile(envFile string) {
	flag.Parse()

	err := godotenv.Load(envFile)
	if err != nil {
		log.Print("Error loading .env file")
	}

	config = appConfig{
		AppName:  os.Getenv("APP_NAME"),
		Host:     os.Getenv("HOST"),
		MongoURI: os.Getenv("MONGO_URI"),
	}
}

func Init() {
	InitWithEnvFile(".env")
}

func GetConfig() appConfig {
	return config
}
