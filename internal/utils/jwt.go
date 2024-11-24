package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// Claims represents the JWT claims.
type Claims struct {
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for the given claims with error handling.
func GenerateJWT(tokenKey string, claims *Claims) (string, error) {
	secretKey := viper.GetString(tokenKey)

	// Check if the secretKey is empty
	if secretKey == "" {
		return "", fmt.Errorf("secret key cannot be empty")
	}

	// Check if the claims are nil
	if claims == nil {
		return "", fmt.Errorf("claims cannot be nil")
	}

	// Check if RegisteredClaims fields are set
	if claims.Issuer == "" {
		return "", fmt.Errorf("issuer cannot be empty")
	}
	if claims.IssuedAt.Time.IsZero() {
		return "", fmt.Errorf("issuedAt cannot be zero")
	}
	if claims.ExpiresAt.Time.IsZero() {
		return "", fmt.Errorf("expiresAt cannot be zero")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %v", err)
	}

	return tokenString, nil
}

// VerifyJWT verifies the JWT token and returns the claims if valid.
func VerifyJWT(tokenKey string, tokenString string) (*Claims, error) {
	secretKey := viper.GetString(tokenKey)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
