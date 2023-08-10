package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

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
	fmt.Printf("%v", config)
}

func Init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	InitWithEnvFile(root + "/.env")
}

func GetConfig() appConfig {
	return config
}
