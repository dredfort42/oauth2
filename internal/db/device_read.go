package db

// DeviceGetUser returns the user associated with the device
func DeviceGetUser(deviceUUID string) (email string, err error) {
	query := `
		SELECT email
		FROM ` + db.tableDevices + `
		WHERE device_uuid = $1;
	`

	err = db.database.QueryRow(query, deviceUUID).Scan(&email)

	return
}
