package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

var myKey = []byte("gin-gorm-oj=key")

// token 签发
func GenerateToken(id, name string) (string, error) {
	UserClaim := &UserClaims{
		Id:   id,
		Name: name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// token 解析
func AnalyseToken(tokenstring string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	Claim, err := jwt.ParseWithClaims(tokenstring, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !Claim.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}
