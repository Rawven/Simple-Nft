package util

import (
	"github.com/dubbogo/gost/log/logger"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

type MyCustomClaims struct {
	UserID int
	jwt.RegisteredClaims
}
type UserInfo struct {
	UserId     int32
	UserName   string
	Address    string
	Balance    int32
	Avatar     string
	PrivateKey string
}

// 签名密钥
const sign_key = "hello jwt"

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func GetJwt(userId int) (string, error) {
	claim := MyCustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                   // 签发者
			Subject:   "Tom",                                           // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(sign_key))
	if err != nil {
		return "", err
	}
	_, err1 := ParseToken(token)
	if err1 != nil {
		return "", err1
	}
	logger.Info(token)
	return token, nil
}

func ParseToken(token_string string) (*int, error) {
	token, _ := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(sign_key), nil //返回签名密钥
	})
	claims, _ := token.Claims.(*MyCustomClaims)
	logger.Info(claims.UserID)
	return &claims.UserID, nil
}
