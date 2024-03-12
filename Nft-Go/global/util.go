package global

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtBlc struct {
	jwt.StandardClaims
}

func GetJwt(key string) (string, error) {
	claims := JwtBlc{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Rawven",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func ParseToken(token string) (*JwtBlc, error) {
	tk, err := jwt.ParseWithClaims(token, &JwtBlc{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Rawven"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tk.Claims.(*JwtBlc)
	if ok && tk.Valid {
		return claims, nil
	}
	return nil, errors.New("token无效")

}
