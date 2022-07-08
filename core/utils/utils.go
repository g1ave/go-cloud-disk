package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/g1ave/go-cloud-disk/core/define"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateNewToke(id uint, name, identity string, second int) (string, error) {
	expiredTime := time.Now().Add(time.Second * time.Duration(second))
	uc := define.UserClaim{
		Id:       id,
		Name:     name,
		Identity: identity,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateCode(codeLength int) (res string) {
	s := "1234567890"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < codeLength; i++ {
		res += string(s[rand.Intn(len(s))])
	}
	return
}
