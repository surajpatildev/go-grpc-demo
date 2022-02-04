package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func Generate() (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
		},
		Username: "slava",
		Role:     "admin",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("my secret"))
}

func main() {
	token, err := Generate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
