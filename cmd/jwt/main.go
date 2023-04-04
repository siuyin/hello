package main

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	pub, priv, err := ed25519.GenerateKey(nil)
	fmt.Println("jwt")
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2000, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(priv)
	fmt.Println(tokenString, err)

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", parsedToken)
	fmt.Printf("%v\n", parsedToken.Valid)
	fmt.Printf("%v\n", parsedToken.Claims)
}
