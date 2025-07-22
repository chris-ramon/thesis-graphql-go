package main

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	handleErr := func(err error) {
		log.Fatal(err)
	}

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"int": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return 20, nil
				},
			},
		},
	})

	implementationSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		handleErr(err)
	}

	result := graphql.Do(graphql.Params{
		Schema:        implementationSchema,
		RequestString: "{int}",
	})

	d, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		handleErr(err)
	}

	log.Printf("%+v", string(d))
}
