package auth

import (
	"errors"
	"fmt"
	"go-graphql-starter/config"
	"go-graphql-starter/person"

	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateToken(tokenString string) (person.Person, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return config.GetJWTSecret(), nil
	})

	if err != nil {
		return person.Person{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		person, _ := claims["data"].(person.Person)
		return person, nil
	}
	return person.Person{}, errors.New("Token not valid")

}
