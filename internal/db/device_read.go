package db

import (
	"database/sql"

	loger "github.com/dredfort42/tools/logprinter"
)

// DeviceExistsCheck checks if a device exists in the database
func DeviceExistsCheck(email string) (result bool) {
	query := `
		SELECT 1
		FROM ` + DB.TableDevices + `
		WHERE email = $1
	`

	err := DB.Database.QueryRow(query, email).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		loger.Error("Failed to check if device exists in the database", err)
	}

	return
}

// DeviceIdentify identifies the device
func DeviceIdentify(deviceUUID, email string) (result bool) {
	query := `
		SELECT 1
		FROM ` + DB.TableDevices + `
		WHERE device_uuid = $1 AND email = $2
	`

	err := DB.Database.QueryRow(query, deviceUUID, email).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		loger.Error("Failed to identify device", err)
	}

	return
}
