package core

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	"tokenutils/services"
)

var jwtSecret = []byte("1122233")
var TokenExpirationTime = 2 * time.Hour

type Claims struct {
	Id int32 `json:"id"`
	jwt.StandardClaims
}

type TokenService struct {
}

func (*TokenService) GetIdByToken(ctx context.Context, req *services.GetIdByTokenRequest, out *services.GetIdByTokenResponse) error {
	token := req.Token
	token = string(token)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) { return jwtSecret, nil })
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			out.UserId = int32(claims.Id)
			return nil
		}
	}
	return err
}

func (*TokenService) GenerateTokenByID(ctx context.Context, req *services.GenerateTokenByIDRequest, out *services.GenerateTokenByIDResponse) error {
	id := req.UserId
	nowTime := time.Now()
	expireTime := nowTime.Add(TokenExpirationTime)
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "1122233",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	out.Token = token
	return err
}

func (*TokenService) VarifyToken(ctx context.Context, req *services.VerifyTokenRequest, out *services.VerifyTokenResponse) error {
	token := req.Token
	// 1 .验证长度
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New("token format error")
	}
	// 2. 验证
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) { return jwtSecret, nil })
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			// 判断当前用户是否匹配
			if claims.Id != req.UserId {
				return errors.New("错误！token对应的id不一致！")
			}
			out.Status = "验证成功！"
			out.Token = token
			return nil
		}
	}
	return err
}
