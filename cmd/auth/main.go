package main

import (
	"auth/internal/api"
	"auth/internal/config"
	"auth/internal/db"
	"os"
	"sort"

	cfg "github.com/dredfort42/tools/configreader"
	loger "github.com/dredfort42/tools/logprinter"
)

func main() {
	err := cfg.GetConfig()
	if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "1" {
		keys := make([]string, 0, len(cfg.Config))
		for key := range cfg.Config {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		loger.Debug("Config:")
		for _, key := range keys {
			value := cfg.Config[key]
			if key == "db.password" {
				value = "********"
			}

			loger.Debug(key, value)
		}
	}

	config.JWTConfigRead()
	config.PasswordConfigRead()

	db.DatabaseInit()
	api.ApiInit()
}
