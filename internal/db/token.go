package db

import (
	s "auth/internal/structs"
)

// TokenAssociatedEmail returns email associated with a token
func TokenAssociatedEmail(id string, token string, tokenType s.TokenType) (email string) {
	var query string

	switch tokenType {
	case s.AccessToken:
		query = `
			SELECT email
			FROM ` + DB.TableSessions + `
			WHERE email = $1 AND access_token = $2
			LIMIT 1;
		`
	case s.RefreshToken:
		query = `
			SELECT email
			FROM ` + DB.TableSessions + `
			WHERE email = $1 AND refresh_token = $2
			LIMIT 1;
		`
	case s.DeviceAccessToken:
		query = `
			SELECT email
			FROM ` + DB.TableDevices + `
			WHERE device_uuid = $1 AND device_access_token = $2
			LIMIT 1;
		`
	case s.DeviceRefreshToken:
		query = `
			SELECT email
			FROM ` + DB.TableDevices + `
			WHERE device_uuid = $1 AND device_refresh_token = $2
			LIMIT 1;
		`
	default:
		return
	}

	DB.Database.QueryRow(query, id, token).Scan(&email)

	return
}

// IsOneTimeRefreshToken checks if a refresh token is a one-time refresh token
func IsOneTimeRefreshToken(refreshToken string) (result bool) {
	query := `
		SELECT 1
		FROM ` + DB.TableSessions + `
		WHERE refresh_token = $1 AND is_one_time = TRUE;
	`

	DB.Database.QueryRow(query, refreshToken).Scan(&result)

	return
}
