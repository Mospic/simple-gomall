package core

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	"tokenutils/service"
	"tokenutils/utils/redis"
)

type TokenService struct {
}

var jwtSecret = []byte("1122233")

type Claims struct {
	Id int32 `json:"id"`
	jwt.StandardClaims
}

func (*TokenService) GetIdByToken(ctx context.Context, req *service.GetIdByTokenRequest, out *service.GetIdByTokenResponse) error {
	token := req.UserToken
	token = string(token)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) { return jwtSecret, nil })
	// TODO store Redis
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			out.UserId = int32(claims.Id)
			return nil

		}
	}
	return err

}

func (*TokenService) StoreTokenToID(ctx context.Context, req *service.StoreTokenToIDRequest, out *emptypb.Empty) error {
	token := req.UserToken
	userID := req.UserId
	token = string(token)
	// 将 Token 和 userID 存储到 Redis，并设置过期时间
	err := redis.RdbJwt.Set(context.Background(), token, userID, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}
