package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("MyNewSecretKey")
func createJWT(username string) {
	fmt.Println("Creating a JWT token.")
	

	token := jwt.NewWithClaims( jwt.SigningMethodHS256 , jwt.MapClaims{
		"sub":username, // Subject (User Identifier)
		"iss": "web_app_go_lang", //Issuer
		"aud": getRole(username), //Auidence (User Role)
		"exp": time.Now().Add(time.Minute * 15).Unix(), //Expiration Time
		"iat": time.Now().Unix(), // Issued At
	})

	fmt.Println("Token headers ",  token.Header)
	fmt.Println("Token Claims are ", token.Claims)

	signedToken, err := token.SignedString(secretKey)
	if err!=nil {
		log.Fatal("Error signing the token", err.Error())
	}
	fmt.Println("Printing the signed token ", signedToken)

	validateJWT(signedToken)
}

func getRole(username string) string {
	if username == "admin" {
		return "admin"
	}
	return "employee"
}

//TODO: Update this to run Both token and Error.
func validateJWT(tokenString string) *jwt.Token {

	claims := jwt.MapClaims{
		"iss": "web_app_go_lang", //Issuer
		"exp": time.Now().Add(time.Minute * 15).Unix(), //Expiration Time
		"iat": time.Now().Unix(), // Issued At
	}

	fmt.Println("Validating JWT Token")
	keyFunc := func(token *jwt.Token) (interface {}, error) {
		return secretKey, nil
	}
	fmt.Println("HS256 name: ", jwt.SigningMethodHS256.Name)
	fmt.Println("HS256 Algo: ", jwt.SigningMethodHS256.Alg())
	token, err := jwt.Parse(tokenString, keyFunc, jwt.WithValidMethods([]string {jwt.SigningMethodHS256.Alg()}) )
	if err != nil {
		log.Fatal("Failed to validate token with error: ", err.Error())
	}

	token , err = jwt.ParseWithClaims(tokenString, claims, keyFunc, jwt.WithValidMethods([]string {jwt.SigningMethodHS256.Name}))

	if err != nil {
		log.Fatal("Failed to validate token with error: ", err.Error())
	}

	fmt.Println("Token validated successfully")
	return token
}