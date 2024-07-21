package db

import (
	"database/sql"

	loger "github.com/dredfort42/tools/logprinter"
)

// DeviceExistsCheck checks if a device exists in the database
func DeviceExistsCheck(email string) (result bool) {
	query := `
		SELECT 1
		FROM ` + db.tableDevices + `
		WHERE email = $1
	`

	err := db.database.QueryRow(query, email).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		loger.Error("Failed to check if device exists in the database", err)
	}

	return
}
