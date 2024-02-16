package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	secret string = "e1t9AuQ79PzVzXbRTUDNDjfOl+KKUIQjsOsOUj7q88A=" //这个未来不能写在代码里，需要写在数据库或者服务器环境变量里
)

// 获取jwt密钥，这个未来不能写在代码里，需要写在数据库或者服务器环境变量里
func GetJwtSecret() string {
	return secret
}

var (
	//可放行的白名单路径
	//whileList = []string{"/ping", "/metrics", "/api/v1/users/login", "/api/v1/users/register", "/ping"}
	whileList = []string{"/user/v1/ping", "/user/v1/metrics", "/user/v1/users/login", "/user/v1/users/register"}
)

// 判断路径是否在白名单中
func isPathInWhiteList(path string) bool {
	for _, v := range whileList {
		if v == path {
			return true
		}
	}
	return false
}

// 创建并签署jwt
func GenerateJWT(userInfo interface{}, secretKey string) (string, error) {
	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userInfo": userInfo,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// 签署 JWT
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
