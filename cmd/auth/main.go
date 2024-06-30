package main

import (
	"auth/internal/api"
	"auth/internal/db"
	"os"

	cfg "github.com/dredfort42/tools/configreader"
	loger "github.com/dredfort42/tools/logprinter"
)

func main() {
	err := cfg.GetConfig()
	if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") == "1" {
		for key, value := range cfg.Config {
			if key == "db.password" {
				value = "********"
			}

			loger.Debug(key, value)
		}
	}

	db.DatabaseInit()
	api.ApiInit()
}
