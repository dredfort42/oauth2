package api

import (
	db "auth/internal/db"
	s "auth/internal/structs"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// getToken generates token
func getToken(id string, expiration int) (response string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)
	tokenClaims["id"] = id
	tokenClaims["exp"] = time.Now().Add(time.Second * time.Duration(expiration)).Unix()

	return token.SignedString([]byte(s.JWTConfig.TokenSecret))
}

// getTokens gets access and refresh tokens
func getTokens(id string, accessTokenExpiration int, refreshTokenExpiration int) (accessToken string, refreshToken string, err error) {
	accessToken, err = getToken(id, accessTokenExpiration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = getToken(id, refreshTokenExpiration)
	if err != nil {
		return "", "", err
	}

	return
}

// parseToken verifies token
func parseToken(token string) (id string, err error) {
	var jwtToken *jwt.Token

	jwtToken, err = jwt.Parse(
		token,
		func(jwtToken *jwt.Token) (interface{}, error) {
			return []byte(s.JWTConfig.TokenSecret), nil
		})
	if err != nil {
		err = errors.New("failed to parse token")
		return
	}

	if time.Now().Unix() > int64(jwtToken.Claims.(jwt.MapClaims)["exp"].(float64)) {
		err = errors.New("token has expired")
		return
	}

	id = jwtToken.Claims.(jwt.MapClaims)["id"].(string)
	if id == "" {
		err = errors.New("failed to get user ID from token")
	}

	return
}

// verifyToken verifies token
func verifyToken(token string, tokenType s.TokenType) (id string, err error) {
	id, err = parseToken(token)
	if err != nil {
		return
	}

	if !db.IsTokenExist(id, token, tokenType) {
		err = errors.New("token does not exist")
	}

	return
}
