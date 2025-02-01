package app

import (
	"os"
)

type AppConfig struct {
	Host string
	Port string
}

func (app *AppConfig) InitService() {
	app.Host = os.Getenv("APP_HOST")
	app.Port = os.Getenv("APP_PORT")
}
