package config

import (
	s "auth/internal/structs"
	"strconv"

	cfg "github.com/dredfort42/tools/configreader"
)

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
