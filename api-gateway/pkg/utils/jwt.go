package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// token密钥 -> 配置类配置？
var jwtSecret = []byte("1122233")

type Claims struct {
	Id int32 `json:"id"`
	jwt.StandardClaims
}

// 签发用户token
func GenerateToken(id int32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour) //有效期为 2 小时
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "DouyinMall_JWT",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) { return jwtSecret, nil })
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token尚未生效")
			} else {
				return nil, errors.New("无法处理该token")
			}
		}
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("无效的token")
}
