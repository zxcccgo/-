package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("是故意的还是不小心的")

func main() {

	token, err := GenerateToken()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(token)
	c, err := ParseToken(token)
	if err != nil {
		log.Println("24", err)
	}
	fmt.Println(c)

}

// 使用jwt传输的信息
type Claim struct {
	Id       int64
	UserName string
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	t := time.Now()
	expireTime := t.Add(300 * time.Second).Unix()

	claims := Claim{
		Id:       10,
		UserName: "zx",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "zx",
			IssuedAt:  time.Now().Unix(),
		},
	}
	fmt.Println(claims)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		log.Println("44")
		panic(err.Error())
	}

	return token, nil
}

func ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Println("57", err)
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {

			return claims, nil
		}
	}
	return nil, err

}
