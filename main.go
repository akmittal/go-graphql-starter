package main

import (
	"fmt"
	"go-graphql-starter/db"
	"go-graphql-starter/schema"
	"log"
	"net/http"

	"github.com/neelance/graphql-go/relay"

	"github.com/neelance/graphql-go"
)

func main() {
	// Connect to DB
	_, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to Database, Exiting...", err)
	}

	parsedSchema, err := graphql.ParseSchema(schema.Schema, &schema.RootResolver{})
	if err != nil {
		fmt.Println("Error parsing schema", err)
	}
	http.Handle("/api", &relay.Handler{Schema: parsedSchema})
	http.HandleFunc("/graphiql", schema.GraphiQLHandler)
	http.ListenAndServe(":8000", nil)
}
