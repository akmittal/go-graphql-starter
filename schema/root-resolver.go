package schema

import (
	"context"
	"fmt"
	"go-graphql-starter/auth"
	"go-graphql-starter/config"
	"go-graphql-starter/db"
	"go-graphql-starter/person"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type RootResolver struct{}

func (resolver *RootResolver) Viewer(ctx context.Context, args struct{ Token string }) (*ViewerResolver, error) {
	user, err := auth.ValidateToken(args.Token)
	if err != nil {
		return nil, err
	}
	return &ViewerResolver{user}, nil
}

func (resolver *RootResolver) Login(ctx context.Context, args struct {
	Username string
	Password string
}) (string, error) {
	db := db.GetConnection()

	var user person.Person

	user = person.Person{Username: args.Username, Password: args.Password}
	personData, err := user.ComparePassword(db)
	fmt.Println(personData)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "go-graphql-starter",
		"iot":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 48).Unix(),
		"data": personData,
	})

	tokenString, err := token.SignedString(config.GetJWTSecret())
	if err != nil {
		return "", err

	}
	return tokenString, nil

}

func (resolver *RootResolver) Signup(ctx context.Context, args struct {
	Person person.PersonInput
}) (person.PersonResolver, error) {
	db := db.GetConnection()

	var user person.Person
	user = person.Person{
		Username:  args.Person.Username,
		Password:  args.Person.Password,
		EmailID:   args.Person.EmailID,
		FirstName: args.Person.FirstName,
		LastName:  args.Person.LastName,
	}

	err := user.Validate()
	if err != nil {
		return person.PersonResolver{}, err
	}
	savedPerson, err := user.Save(db)
	if err != nil {
		return person.PersonResolver{}, err
	}
	return person.PersonResolver{Person: savedPerson}, nil

}
