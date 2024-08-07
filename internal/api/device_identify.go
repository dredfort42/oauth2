package api

import (
	s "auth/internal/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeviceIdentify identifies the device
func DeviceIdentify(c *gin.Context) {
	var errorResponse s.ResponseError
	var err error

	clientID := c.Request.URL.Query().Get("client_id")
	if clientID == "" {
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "client_id is required"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	accessToken := c.GetHeader("Authorization")
	if len(accessToken) > 50 && accessToken[:7] == "Bearer " {
		accessToken = accessToken[7:]
	} else {
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "missing device access token"
		c.IndentedJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	email, err := verifyToken(accessToken, s.DeviceAccessToken)
	if err != nil {
		errorResponse.Error = "token_error"
		errorResponse.ErrorDescription = "failed to verify device access token"
		c.IndentedJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": email})
}
