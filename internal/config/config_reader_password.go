package config

import (
	s "auth/internal/structs"
	"strconv"

	cfg "github.com/dredfort42/tools/configreader"
)

// PasswordConfigRead reads password configuration
func PasswordConfigRead() {
	passwordMinLength, _ := strconv.Atoi(cfg.Config["password.require.length"])

	s.PasswordConfig = s.PasswordParamethers{
		RequireUpper:  cfg.Config["password.require.upper"] == "true",
		RequireLower:  cfg.Config["password.require.lower"] == "true",
		RequireNumber: cfg.Config["password.require.number"] == "true",
		RequireSymbol: cfg.Config["password.require.symbol"] == "true",
		RequireLength: passwordMinLength,
	}
}
