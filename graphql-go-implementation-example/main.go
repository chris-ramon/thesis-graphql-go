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

	enumType := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Enum",
		Description: "A enum.",
		Values: graphql.EnumValueConfigMap{
			"FIRST_ENUM": &graphql.EnumValueConfig{
				Value:       1,
				Description: "First enum value.",
			},
			"SECOND_ENUM": &graphql.EnumValueConfig{
				Value:       2,
				Description: "Second enum value.",
			},
		},
	})

	userInputType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        "UserInput",
		Description: "Input type for user data.",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "The name of the user.",
			},
			"email": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "The email of the user.",
			},
			"age": &graphql.InputObjectFieldConfig{
				Type:        graphql.Int,
				Description: "The age of the user.",
			},
		},
	})

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
	var searchResultUnion *graphql.Union

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
				switch obj["type"] {
				case "user":
					return userType
				case "product":
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

	userTypeNonNull := graphql.NewObject(graphql.ObjectConfig{
		Name:        "UserNonNull",
		Description: "A user with non-null fields.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "The non-null ID of the user.",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The non-null name of the user.",
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

	productTypeNonNull := graphql.NewObject(graphql.ObjectConfig{
		Name:        "ProductNonNull",
		Description: "A product with non-null fields.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "The non-null ID of the product.",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The non-null name of the product.",
			},
			"price": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Float),
				Description: "The non-null price of the product.",
			},
		},
	})

	searchResultUnion = graphql.NewUnion(graphql.UnionConfig{
		Name:        "SearchResult",
		Description: "A union of User and Product types.",
		Types:       []*graphql.Object{userType, productType},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if obj, ok := p.Value.(map[string]interface{}); ok {
				switch obj["type"] {
				case "user":
					return userType
				case "product":
					return productType
				}
			}
			return nil
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
			"enum": &graphql.Field{
				Type: enumType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return 2, nil
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
			"searchResult": &graphql.Field{
				Type: searchResultUnion,
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Description: "type of the search result",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					searchType := p.Args["type"].(string)
					if searchType == "user" {
						return map[string]interface{}{
							"type": "user",
							"id":   "user-1",
							"name": "John Doe",
						}, nil
					}
					if searchType == "product" {
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
			"createUser": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Description: "input for creating a user",
						Type:        userInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					return fmt.Sprintf("User created with name: %v, email: %v, age: %v", input["name"], input["email"], input["age"]), nil
				},
			},
			"stringList": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return []string{"first string", "second string", "third string"}, nil
				},
			},
			"objectList": &graphql.Field{
				Type: graphql.NewList(objectType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return []interface{}{
						map[string]interface{}{"name": "First object in list"},
						map[string]interface{}{"name": "Second object in list"},
						map[string]interface{}{"name": "Third object in list"},
					}, nil
				},
			},
			"userNonNull": &graphql.Field{
				Type: userTypeNonNull,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{
						"id":   "user-non-null-1",
						"name": "John Doe Non-Null",
					}, nil
				},
			},
			"productNonNull": &graphql.Field{
				Type: productTypeNonNull,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return map[string]interface{}{
						"id":    "product-non-null-1",
						"name":  "GraphQL Book Non-Null",
						"price": 39.99,
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
      enum
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
      searchResult(type: "user") {
        ... on User {
          id
          name
        }
        ... on Product {
          id
          name
          price
        }
      }
      createUser(input: { name: "Alice", email: "alice@example.com", age: 30 })
      stringList
      objectList {
        name
      }
      userNonNull {
        id
        name
      }
      productNonNull {
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
