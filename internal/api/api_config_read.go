package api

import (
	s "auth/internal/structs"
	"strconv"

	cfg "github.com/dredfort42/tools/configreader"
)

// ServiceConfigRead reads service configuration
func ServiceConfigRead() {
	host = cfg.Config["auth.host"]
	if host == "" {
		panic("auth.host is not set")
	}

	port = cfg.Config["auth.port"]
	if port == "" {
		panic("auth.port is not set")
	}

	corsStatus = cfg.Config["auth.cors"]
	if corsStatus == "" {
		panic("auth.cors is not set")
	}

	deviceVerificationURI = cfg.Config["auth.device.verification.url"]
	if deviceVerificationURI == "" {
		panic("auth.device.verification.url is not set")
	}

	deviceVerificationCodeCharSet = cfg.Config["auth.device.verification.code.charset"]
	if deviceVerificationCodeCharSet == "" {
		panic("auth.device.verification.code.charset is not set")
	}

	var err error

	deviceVerificationCodeLength, err = strconv.Atoi(cfg.Config["auth.device.verification.code.length"])
	if err != nil {
		panic("auth.device.verification.code.length is not set")
	}

	deviceVerificationCodeExpiration, err = strconv.Atoi(cfg.Config["auth.device.verification.code.expiration"])
	if err != nil {
		panic("auth.device.verification.code.expiration is not set")
	}

	deviceVerificationCodeAttempts, err = strconv.Atoi(cfg.Config["auth.device.verification.code.attempts"])
	if err != nil {
		panic("auth.device.verification.code.attempts is not set")
	}
}

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

// JWTConfigRead reads JWT configuration
func JWTConfigRead() {
	s.JWTConfig = s.JWTParamethers{
		TokenSecret: cfg.Config["jwt.secret"],
	}
	if s.JWTConfig.TokenSecret == "" {
		panic("JWT secret is not set")
	}

	var expiration int
	var err error

	expiration, err = strconv.Atoi(cfg.Config["jwt.onetime.access.token.expiration"])
	if err != nil {
		panic("JWT onetime access token expiration is not set")
	}
	s.JWTConfig.OneTimeAccessTokenExpiration = expiration

	expiration, err = strconv.Atoi(cfg.Config["jwt.onetime.refresh.token.expiration"])
	if err != nil {
		panic("JWT onetime refresh token expiration is not set")
	}
	s.JWTConfig.OneTimeRefreshTokenExpiration = expiration

	expiration, err = strconv.Atoi(cfg.Config["jwt.browser.access.token.expiration"])
	if err != nil {
		panic("JWT browser access token expiration is not set")
	}
	s.JWTConfig.BrowserAccessTokenExpiration = expiration

	expiration, err = strconv.Atoi(cfg.Config["jwt.browser.refresh.token.expiration"])
	if err != nil {
		panic("JWT browser refresh token expiration is not set")
	}
	s.JWTConfig.BrowserRefreshTokenExpiration = expiration

	expiration, err = strconv.Atoi(cfg.Config["jwt.device.access.token.expiration"])
	if err != nil {
		panic("JWT device access token expiration is not set")
	}
	s.JWTConfig.DeviceAccessTokenExpiration = expiration

	expiration, err = strconv.Atoi(cfg.Config["jwt.device.refresh.token.expiration"])
	if err != nil {
		panic("JWT device refresh token expiration is not set")
	}
	s.JWTConfig.DeviceRefreshTokenExpiration = expiration
}
