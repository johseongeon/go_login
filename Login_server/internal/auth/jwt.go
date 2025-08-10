package auth

import (
	"main/internal/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Username string `json:"username"`
	User_id  string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(loginUser *user.User, skey []byte) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &CustomClaims{
		Username: loginUser.Username,
		User_id:  loginUser.User_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tapvpn",
		},
	}
	// hashing algorithm : HS256(HMAC + SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token == nil {
		return "", jwt.ErrTokenMalformed
	}

	tokenString, err := token.SignedString(skey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
