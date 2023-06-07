package commons

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("夏天夏天悄悄过去留下小秘密")

type MyClaims struct {
	UserId int
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(userId int) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "matuto-blog",
			Subject:   "user token",
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	token, err := claims.SignedString(MySecret)
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	// 解析Tokne
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		log.Printf("ParseToken, parseToken is error: %v", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
