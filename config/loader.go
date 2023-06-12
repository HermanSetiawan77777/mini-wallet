package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "wallet-julo"

func GetAppRootDirectory() string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}

func LoadEnv() {
	rootPath := GetAppRootDirectory()

	err := godotenv.Load(rootPath + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
