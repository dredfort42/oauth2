package api

import (
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
