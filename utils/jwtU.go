package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type AdminClaims struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var adminKey = []byte("hnggyyds")

// GenerateToken
// 生成 token

func GenerateAdminToken(name string, password string) string {
	AdminClaim := &AdminClaims{
		Name:           name,
		PassWord:       password,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AdminClaim)
	tokenString, err := token.SignedString(adminKey)
	if err != nil {
		return ""
	}
	return tokenString
}

// AnalyseToken
// 解析 token

func AnalyseAdminToken(tokenString string) (*AdminClaims, error) {
	adminClaim := new(AdminClaims)
	claims, err := jwt.ParseWithClaims(tokenString, adminClaim, func(token *jwt.Token) (interface{}, error) {
		return adminKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return adminClaim, nil
}
