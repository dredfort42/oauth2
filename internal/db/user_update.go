package db

import loger "github.com/dredfort42/tools/logprinter"

// UserPasswordChange updates a user's password in the database
func UserPasswordChange(email string, newPassword string) (err error) {
	query := `
		UPDATE ` + DB.TableUsers + `
		SET	password_hash = crypt($2, gen_salt('bf')),
			updated_at = CURRENT_TIMESTAMP
		WHERE email = $1;
	`

	_, err = DB.Database.Exec(query, email, newPassword)
	if err != nil {
		loger.Error("Failed to update user in the database", err)
	}

	return
}

// UserEmailChange updates a user's email address in the database
func UserEmailChange(email string, newEmail string) (err error) {
	query := `
		UPDATE ` + DB.TableUsers + `
		SET	email = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE email = $1;
	`

	_, err = DB.Database.Exec(query, email, newEmail)
	if err != nil {
		loger.Error("Failed to update user in the users table", err)
	}

	if !DeviceExistsCheck(email) {
		return
	}

	query = `
		UPDATE ` + DB.TableDevices + `
		SET	email = $2
		WHERE email = $1;
	`

	_, err = DB.Database.Exec(query, email, newEmail)
	if err != nil {
		loger.Error("Failed to update devices in the devices table", err)
	}

	return
}
