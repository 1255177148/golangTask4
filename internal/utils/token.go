package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var accessSecret = []byte("access_secret")
var refreshSecret = []byte("refresh_secret")

// GenerateAccessToken 生成 access token
func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"typ":     "access",
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(accessSecret)
}

// GenerateRefreshToken 生成refresh token
func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"typ":     "refresh",
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(refreshSecret)
}

// ParseToken 从token中取出user id
func ParseToken(tokenStr string, isRefresh bool) (uint, error) {
	secret := accessSecret
	if isRefresh {
		secret = refreshSecret
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return uint(claims["user_id"].(float64)), nil
}

// Sha256Hex 计算字符串的 SHA-256 哈希，返回十六进制字符串
func Sha256Hex(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
