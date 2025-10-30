package auth

import (
	"errors"
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

/* Token struct declared in jwt package
type Token struct {
	Raw       string         // Raw contains the raw token.  Populated when you [Parse] a token
	Method    SigningMethod  // Method is the signing method used or to be used
	Header    map[string]any // Header is the first segment of the token in decoded form
	Claims    Claims         // Claims is the second segment of the token in decoded form
	Signature []byte         // Signature is the third segment of the token in decoded form.  Populated when you Parse a token
	Valid     bool           // Valid specifies if the token is valid.  Populated when you Parse/Verify a token -> using this field to check if the token is valid
}

// ParseWithClaims parses a token and returns the parsed token with claims.
// The provided claims must be a pointer to a struct that implements the Claims interface.
// If the token is valid, the claims will be populated with the token's claims.
*/

func ValidateJWT(tokenString, jwtSecret string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	// Parse and validate
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Optional: Check issuer (defensive)
	if claims.Issuer != "tapvpn" {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}
