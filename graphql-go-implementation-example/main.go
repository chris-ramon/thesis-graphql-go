package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	handleErr := func(err error) {
		log.Fatal(err)
	}

	objectType := graphql.NewObject(graphql.ObjectConfig{
		Name: "object",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Description: "The name of the object.",
				Type:        graphql.String,
			},
		},
	})

	objectTypeWithArguments := graphql.NewObject(graphql.ObjectConfig{
		Name: "objectTypeWithArguments",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Description: "The name of the object.",
				Type:        graphql.String,
			},
		},
	})

	var nodeInterface *graphql.Interface
	var userType *graphql.Object
	var productType *graphql.Object

	nodeInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Node",
		Description: "An object with an ID.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The ID of the object.",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if obj, ok := p.Value.(map[string]interface{}); ok {
				if objType, exists := obj["type"]; exists && objType == "user" {
					return userType
				}
				if objType, exists := obj["type"]; exists && objType == "product" {
					return productType
				}
			}
			return nil
		},
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A user.",
		Interfaces:  []*graphql.Interface{nodeInterface},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The ID of the user.",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user.",
			},
		},
	})

	productType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Product",
		Description: "A product.",
		Interfaces:  []*graphql.Interface{nodeInterface},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The ID of the product.",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the product.",
			},
			"price": &graphql.Field{
				Type:        graphql.Float,
				Description: "The price of the product.",
			},
		},
	})

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
			"object": &graphql.Field{
				Type: objectType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj := struct {
						Name string
					}{
						Name: "Name of the object instance.",
					}
					return obj, nil
				},
			},
			"objectWithArguments": &graphql.Field{
				Type: objectTypeWithArguments,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the object with arguments",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					obj := struct {
						Name string
					}{
						Name: fmt.Sprintf("Name of the object with arguments instance, id: %v", p.Args["id"]),
					}
					return obj, nil
				},
			},
			"node": &graphql.Field{
				Type: nodeInterface,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the node",
						Type:        graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					if id == "user-1" {
						return map[string]interface{}{
							"type": "user",
							"id":   "user-1",
							"name": "John Doe",
						}, nil
					}
					if id == "product-1" {
						return map[string]interface{}{
							"type":  "product",
							"id":    "product-1",
							"name":  "GraphQL Book",
							"price": 29.99,
						}, nil
					}
					return nil, nil
				},
			},
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{
						"type": "user",
						"id":   "user-1",
						"name": "John Doe",
					}, nil
				},
			},
			"product": &graphql.Field{
				Type: productType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{
						"type":  "product",
						"id":    "product-1",
						"name":  "GraphQL Book",
						"price": 29.99,
					}, nil
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
      object {
        name
      }
      objectWithArguments(id: "1") {
        name
      }
      node(id: "user-1") {
        id
        ... on User {
          name
        }
        ... on Product {
          name
          price
        }
      }
      user {
        id
        name
      }
      product {
        id
        name
        price
      }
    }
    `,
	})

	d, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		handleErr(err)
	}

	log.Printf("%+v", string(d))
}
