// This package is a demonstration how to build and use a GraphQL server in Go
package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	// We create yet another Fields map, one which holds all the different queries
	fields := graphql.Fields{
		// We define the Gophers query
		"gophers": &graphql.Field{
			// It a String, FOR NOW
			Type: graphql.String,
			// Resolve is the function used to look up data
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "hello", nil
			},
			// Description explains the field
			Description: "Query all Gophers",
		},
	}
	// Create the Root Query that is used to start each query
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// Now combine all Objects into a Schema Configuration
	schemaConfig := graphql.SchemaConfig{
		// Query is the root object query schema
		Query: graphql.NewObject(rootQuery)}
	// Create a new GraphQL Schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	StartServer(&schema)
}

// StartServer will trigger the server with a Playground
func StartServer(schema *graphql.Schema) {
	// Create a new HTTP handler
	h := handler.New(&handler.Config{
		Schema: schema,
		// Pretty print JSON response
		Pretty: true,
		// Host a GraphiQL Playground to use for testing Queries
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
