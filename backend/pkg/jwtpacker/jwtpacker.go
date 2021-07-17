package jwtpacker

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"backend/config"
)

var (
	TokenExp        = time.Hour * 24
	RefreshTokenExp = time.Hour * 24
)

func GenerateURLSafeTimedToken(claims jwt.MapClaims) (string, error) {
	claims["tokenExp"] = time.Now().Add(TokenExp).Unix()
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tok.SignedString([]byte(config.Cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (bool, jwt.MapClaims) {
	var claims jwt.MapClaims
	tok, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (i interface{}, err error) {
		return []byte(config.Cfg.JWTSecret), nil
	})
	if err != nil {
		return false, nil
	}

	if !tok.Valid {
		return false, nil
	}

	return true, claims
}

func GenerateLoginToken(claims jwt.MapClaims) (string, string, error) {
	claims["tokenExp"] = time.Now().Add(TokenExp).Unix()
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tok.Header["kid"] = "signin_1"

	token, err := tok.SignedString([]byte(config.Cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	claims["tokenExp"] = time.Now().Add(RefreshTokenExp).Unix()
	rfToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rfToken.Header["kid"] = "signin_2"

	refreshToken, err := rfToken.SignedString([]byte(config.Cfg.JWTRefreshSecret))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
