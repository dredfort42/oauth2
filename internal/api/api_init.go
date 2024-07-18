package api

import (
	"os"

	loger "github.com/dredfort42/tools/logprinter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var host string
var port string
var corsStatus string
var deviceVerificationURI string
var deviceVerificationCodeCharSet string
var deviceVerificationCodeLength int
var deviceVerificationCodeExpiration int
var deviceVerificationCodeAttempts int

// ApiInit starts the web service
func ApiInit() {
	ServiceConfigRead()

	if os.Getenv("DEBUG") != "true" && os.Getenv("DEBUG") != "1" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	if corsStatus == "true" || corsStatus == "1" {
		router.Use(cors.Default())
	}

	router.POST("/api/v1/auth/user/signup", UserSignUp)
	router.GET("/api/v1/auth/user/identify", UserIdentify)
	router.DELETE("/api/v1/auth/user/delete", UserDelete)
	router.POST("/api/v1/auth/user/login", UserLogIn)
	router.POST("/api/v1/auth/user/refresh", UserRefresh)
	router.POST("/api/v1/auth/user/logout", UserLogOut)
	router.POST("/api/v1/auth/user/password", UserPasswordChange)
	router.POST("/api/v1/auth/user/email", UserEmailChange)
	router.POST("/api/v1/auth/device/authorize", DeviceAuthorize)
	router.POST("/api/v1/auth/device/verify", DeviceVerify)
	router.POST("/api/v1/auth/device/token", DeviceTokens)
	router.GET("/api/v1/auth/device/identify", DeviceIdentify)
	router.DELETE("/api/v1/auth/device/delete", DeviceDelete)
	router.POST("/api/v1/auth/device/refresh", DeviceRefresh)

	url := host + ":" + port
	loger.Success("Service successfully started", url)
	router.Run(url)
}
