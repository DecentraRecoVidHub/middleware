package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
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

/*
验证jwt是否合法，是否过期
实现思路：
这个 JWT 中间件函数接收一个 secretKey 参数作为 JWT 密钥，并返回一个 gin.HandlerFunc 函数。
在 gin.HandlerFunc 函数中，首先从请求头中获取 JWT Token，然后对 Token 进行解析和验证，
如果 Token 无效，则返回 401 Unauthorized 错误并终止请求。
如果 Token 验证通过，则调用 c.Next() 使请求继续执行。
*/
func JWTAuth(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if isPathInWhiteList(path) {
			c.Next()
			return
		}
		// 从请求头中获取 JWT Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			FailUnauthorized(c, nil, "Missing Authorization Header")
			c.Abort()
			return
		}
		authToken := strings.Split(authHeader, ".")
		fmt.Println(len(authToken))
		//校验jwt是否由3个部分构成
		if len(authToken) != 3 {
			FailUnauthorized(c, nil, "Invalid Authorization Header")
			c.Abort()
			return
		}

		// 验证 JWT Token 是否合法
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			FailUnauthorized(c, nil, "Invalid Token:Token signature verification failed.")
			c.Abort()
			return
		}

		// 检查 JWT Token 是否已经过期
		claims, ok := token.Claims.(jwt.MapClaims)

		// 获取UserRole字段的值
		userRole, _ := claims["userInfo"].(map[string]interface{})["userRole"].(string)

		if !ok {
			FailUnauthorized(c, nil, "Invalid Token:Token format is invalid")
			c.Abort()
			return
		}
		exp, ok := claims["exp"].(float64)
		if !ok {
			FailUnauthorized(c, nil, "Invalid Token:Token format is invalid,not float64")
			c.Abort()
			return
		}
		if int64(exp) < time.Now().Unix() {
			FailUnauthorized(c, nil, "Token Expired:The token has expired.")
			c.Abort()
			return
		}

		// 将UserRole传递给处理函数
		c.Set("userRole", userRole)

		// 请求继续执行
		c.Next()
	}
}
