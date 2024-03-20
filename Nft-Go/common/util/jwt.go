package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtBlc struct {
	userId int32
	jwt.StandardClaims
}
type UserInfo struct {
	UserId     int32
	UserName   string
	Address    string
	Balance    int32
	Avatar     string
	PrivateKey string
}

func GetJwt(key string, userId int32) (string, error) {
	claims := JwtBlc{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Rawven",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func ParseToken(token string) (*int32, error) {
	tk, err := jwt.ParseWithClaims(token, &JwtBlc{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Rawven"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tk.Claims.(*JwtBlc)
	if ok && tk.Valid {
		return &claims.userId, nil
	}
	return nil, errors.New("token无效")
}
