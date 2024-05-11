package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type JWT struct {
	UserUUID    *uuid.UUID
	AccessLevel AccessLevel
	Permissions map[string]interface{}
}

func MarshalJWT(secret string, userUuid string, accessLevel AccessLevel, permissions map[string]bool) (string, error) {
	// Create a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":          userUuid,
		"iat":          time.Now().Unix(),
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
		"access_level": accessLevel,
		"permissions":  permissions,
	})

	// Sign the JWT using the secret key
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed signing jwt: %w", err)
	}

	return signedToken, nil
}

func UnmarshalJWTFromHeader(secret, authHeader string) (*JWT, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("missing Authorization header")
	}

	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) != 2 {
		return nil, fmt.Errorf("invalid Authorization header, expecting Bearer <token>")
	}

	if authSplit[0] != "Bearer" {
		return nil, fmt.Errorf("invalid Authorization header, only Bearer is supported")
	}

	// Extract the JWT from the header
	jwtString := authSplit[1]

	return UnmarshalJWT(secret, jwtString)
}

func UnmarshalJWT(secret, jwtString string) (*JWT, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the alg is what's expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed parsing jwt")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed parsing claims")
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	expiration := time.Unix(int64(claims["exp"].(float64)), 0)
	if expiration.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}

	uuidStr := claims["sub"].(string)
	var userUUID *uuid.UUID
	if uuidStr != "" {
		uuid, err := uuid.Parse(uuidStr)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid uuid")
		}

		userUUID = &uuid
	}

	return &JWT{
		UserUUID:    userUUID,
		AccessLevel: AccessLevel(claims["access_level"].(float64)),
		Permissions: claims["permissions"].(map[string]interface{}),
	}, nil
}
