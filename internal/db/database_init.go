package db

import (
	"database/sql"

	cfg "github.com/dredfort42/tools/configreader"
	loger "github.com/dredfort42/tools/logprinter"
	_ "github.com/lib/pq"
)

// Database is the database struct
type Database struct {
	Database      *sql.DB
	TableUsers    string
	TableSessions string
	TableDevices  string
}

var DB Database

// DatabaseInit initializes the database
func DatabaseInit() {
	DB.TableUsers = cfg.Config["db.table.users"]
	if DB.TableUsers == "" {
		panic("Table users is not set")
	}

	DB.TableSessions = cfg.Config["db.table.sessions"]
	if DB.TableSessions == "" {
		panic("Table sessions is not set")
	}

	DB.TableDevices = cfg.Config["db.table.devices"]
	if DB.TableDevices == "" {
		panic("Table devices is not set")
	}

	databaseConnect()
	tablesCheck()
	databaseCleanerStart()

	loger.Success("Database successfully initialized")
}
