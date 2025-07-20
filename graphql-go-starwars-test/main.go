package main

import (
	"github.com/graphql-go/graphql/testutil"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:     &testutil.StarWarsSchema,
		Pretty:     false,
		GraphiQL:   true,
		Playground: false,
	})

	log.Println("running on por :8080")
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
