package api

import (
	"auth/internal/db"
	s "auth/internal/structs"
	"encoding/base64"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// IsUserPasswordStrength checks the password strength
func IsUserPasswordStrength(password string) (result bool) {
	if len(password) < s.PasswordConfig.RequireLength {
		return
	}

	var hasUpper bool
	var hasLower bool
	var hasNumber bool
	var hasSymbol bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			hasSymbol = true
		}
	}

	if (s.PasswordConfig.RequireUpper && !hasUpper) ||
		(s.PasswordConfig.RequireLower && !hasLower) ||
		(s.PasswordConfig.RequireNumber && !hasNumber) ||
		(s.PasswordConfig.RequireSymbol && !hasSymbol) {
		return
	}

	return true
}

// IsUserPasswordContainsForbiddenCharacters checks the password for forbidden symbols
func IsUserPasswordContainsForbiddenCharacters(password string) (result bool) {
	forbiddenSymbols := " <>;&%\x00\r\n\\"

	for _, char := range password {
		if strings.ContainsRune(forbiddenSymbols, char) {
			return true
		}
	}

	return
}

// UserPasswordChange changes the user password
func UserPasswordChange(c *gin.Context) {
	var changePasswordRequest s.UserChangePasswordRequest
	var errorResponse s.ResponseError
	var err error

	err = c.BindJSON(&changePasswordRequest)
	if err != nil {
		errorResponse.Error = "invalid_request"
		errorResponse.ErrorDescription = "invalid JSON in the request body | " + err.Error()
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if changePasswordRequest.OldPassword == "" || changePasswordRequest.NewPassword == "" {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "old password and new password are required"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if changePasswordRequest.OldPassword == changePasswordRequest.NewPassword {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "old password and new password are the same"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	oldPassword, err := base64.StdEncoding.DecodeString(changePasswordRequest.OldPassword)
	if err != nil {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "old password is invalid"
		c.IndentedJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	newPassword, err := base64.StdEncoding.DecodeString(changePasswordRequest.NewPassword)
	if err != nil {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "new password is invalid"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	changePasswordRequest.OldPassword = string(oldPassword)
	changePasswordRequest.NewPassword = string(newPassword)

	if IsUserPasswordContainsForbiddenCharacters(changePasswordRequest.NewPassword) {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "password contains forbidden characters"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	if !IsUserPasswordStrength(changePasswordRequest.NewPassword) {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "password isn't strong enough"
		c.IndentedJSON(http.StatusBadRequest, errorResponse)
		return
	}

	accessToken, err := c.Cookie("access_token")
	if err != nil {
		errorResponse.Error = "token_error"
		errorResponse.ErrorDescription = "missing access token"
		c.IndentedJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	email, err := verifyToken(accessToken, s.AccessToken)
	if err != nil {
		errorResponse.Error = "token_error"
		errorResponse.ErrorDescription = "failed to verify access token"
		c.IndentedJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	if !db.IsUserPasswordCorrect(email, changePasswordRequest.OldPassword) {
		errorResponse.Error = "invalid_parameter"
		errorResponse.ErrorDescription = "old password is invalid"
		c.IndentedJSON(http.StatusUnauthorized, errorResponse)
		return
	}

	err = db.UserPasswordChange(email, changePasswordRequest.NewPassword)
	if err != nil {
		errorResponse.Error = "database_error"
		errorResponse.ErrorDescription = "failed to update user password in the database | " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user password successfully changed"})
}
