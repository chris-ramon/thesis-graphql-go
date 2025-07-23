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
			"float": &graphql.Field{
				Type: graphql.Float,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return 20.01, nil
				},
			},
			"string": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "str", nil
				},
			},
			"boolean": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return true, nil
				},
			},
			"ID": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "d983b9d9-681c-4059-b5a3-5329d1c6f82d", nil
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
		Schema: implementationSchema,
		RequestString: `
    {
      int
      float
      string
      boolean
      ID
    }
    `,
	})

	d, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		handleErr(err)
	}

	log.Printf("%+v", string(d))
}
