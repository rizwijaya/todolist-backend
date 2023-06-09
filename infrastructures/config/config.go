package config

import (
	"os"
)

func New() LoadConfig {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }
	return LoadConfig{
		App: App{
			Mode: os.Getenv("APP_MODE"),
			Name: os.Getenv("APP_NAME"),
			Port: os.Getenv("APP_PORT"),
			Url:  os.Getenv("APP_URL"),
		},
		Database: Database{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Dbname:   os.Getenv("MYSQL_DBNAME"),
		},
	}
}
