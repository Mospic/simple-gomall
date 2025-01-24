package core

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
	"tokenutils/service"
)

var jwtSecret = []byte("1122233")
var TokenExpirationTime = 2 * time.Hour

type Claims struct {
	Id int32 `json:"id"`
	jwt.StandardClaims
}

type TokenService struct {
}

func (*TokenService) GetIdByToken(ctx context.Context, req *service.GetIdByTokenRequest, out *service.GetIdByTokenResponse) error {
	token := req.UserToken
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

func (*TokenService) GenerateTokenByID(ctx context.Context, req *service.GenerateTokenByIDRequest, out *service.GenerateTokenByIDResponse) error {
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
	out.UserToken = token
	return err
}
