package util

import (
	"github.com/dubbogo/gost/log/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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
var key string

// 随机字符串
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func initJwt() {
	key = viper.GetString("key")
}
func randStr(str int) string {
	bytes := make([]rune, str)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
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
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(key))
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

func ParseToken(tokenStr string) (*int, error) {
	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil //返回签名密钥
	})
	claims, _ := token.Claims.(*MyCustomClaims)
	logger.Info(claims.UserID)
	return &claims.UserID, nil
}
