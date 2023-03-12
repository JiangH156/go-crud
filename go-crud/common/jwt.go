package common

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaims struct {
	UserId uint
	jwt.RegisteredClaims
}

// JWT过期时间
const TokenExpireDuration = time.Hour * 24

// 用于签名的字符串
var jwtScret = []byte("a_secret_create")

// 生成JWT
func GenToken(userId uint) (string, error) {
	//创建一个自己的声明
	claims := UserClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "john",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	//使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(jwtScret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*UserClaims, error) {
	// 自定义claims使用
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtScret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("权限不够")
}
