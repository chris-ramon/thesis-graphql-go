package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	productInputType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        "ProductInput",
		Description: "Input type for product data.",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.String,
				Description: "The name of the product.",
			},
			"price": &graphql.InputObjectFieldConfig{
				Type:        graphql.Float,
				Description: "The price of the product.",
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

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Description: "input for creating a user",
						Type:        userInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					return map[string]interface{}{
						"type": "user",
						"id":   fmt.Sprintf("user-%d", time.Now().Unix()),
						"name": input["name"],
					}, nil
				},
			},
			"createProduct": &graphql.Field{
				Type: productType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Description: "input for creating a product",
						Type:        productInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					return map[string]interface{}{
						"type":  "product",
						"id":    fmt.Sprintf("product-%d", time.Now().Unix()),
						"name":  input["name"],
						"price": input["price"],
					}, nil
				},
			},
			"updateUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the user to update",
						Type:        graphql.NewNonNull(graphql.ID),
					},
					"input": &graphql.ArgumentConfig{
						Description: "input for updating a user",
						Type:        userInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					input := p.Args["input"].(map[string]interface{})
					return map[string]interface{}{
						"type": "user",
						"id":   id,
						"name": input["name"],
					}, nil
				},
			},
			"updateProduct": &graphql.Field{
				Type: productType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the product to update",
						Type:        graphql.NewNonNull(graphql.ID),
					},
					"input": &graphql.ArgumentConfig{
						Description: "input for updating a product",
						Type:        productInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					input := p.Args["input"].(map[string]interface{})
					return map[string]interface{}{
						"type":  "product",
						"id":    id,
						"name":  input["name"],
						"price": input["price"],
					}, nil
				},
			},
			"deleteUser": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the user to delete",
						Type:        graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					return fmt.Sprintf("User with id: %s deleted successfully", id), nil
				},
			},
			"deleteProduct": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the product to delete",
						Type:        graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					return fmt.Sprintf("Product with id: %s deleted successfully", id), nil
				},
			},
		},
	})

	implementationSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		handleErr(err)
	}

	queryResult := graphql.Do(graphql.Params{
		Schema: implementationSchema,
		RequestString: `
    query ExampleQuery($skipUserName: Boolean!, $skipProductPrice: Boolean!, $includeUserName: Boolean!, $includeProductPrice: Boolean!) {
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
          name @skip(if: $skipUserName) @include(if: $includeUserName)
        }
        ... on Product {
          name
          price @skip(if: $skipProductPrice) @include(if: $includeProductPrice)
        }
      }
      user {
        id
        name @skip(if: $skipUserName) @include(if: $includeUserName)
      }
      product {
        id
        name
        price @skip(if: $skipProductPrice) @include(if: $includeProductPrice)
      }
      searchResult(type: "user") {
        ... on User {
          id
          name @skip(if: $skipUserName) @include(if: $includeUserName)
        }
        ... on Product {
          id
          name
          price @skip(if: $skipProductPrice) @include(if: $includeProductPrice)
        }
      }
      stringList
      objectList {
        name
      }
      userNonNull {
        id
        name @skip(if: $skipUserName) @include(if: $includeUserName)
      }
      productNonNull {
        id
        name
        price @skip(if: $skipProductPrice) @include(if: $includeProductPrice)
      }
    }
    `,
		VariableValues: map[string]interface{}{
			"skipUserName":     false,
			"skipProductPrice": true,
			"includeUserName":     true,
			"includeProductPrice": false,
		},
		OperationName: "ExampleQuery",
	})

	d, err := json.MarshalIndent(queryResult, "", "    ")
	if err != nil {
		handleErr(err)
	}

	log.Printf("Query results: %+v", string(d))

	// Now run the mutation
	mutationResult := graphql.Do(graphql.Params{
		Schema: implementationSchema,
		RequestString: `
    mutation ExampleMutation {
      createUser(input: { name: "Alice", email: "alice@example.com", age: 30 }) {
        id
        name
      }
      createProduct(input: { name: "GraphQL Guide", price: 49.99 }) {
        id
        name
        price
      }
      updateUser(id: "user-1", input: { name: "Alice Updated", email: "alice.updated@example.com", age: 31 }) {
        id
        name
      }
      updateProduct(id: "product-1", input: { name: "GraphQL Guide Updated", price: 59.99 }) {
        id
        name
        price
      }
      deleteUser(id: "user-2")
      deleteProduct(id: "product-2")
    }
    `,
		OperationName: "ExampleMutation",
	})

	d2, err := json.MarshalIndent(mutationResult, "", "    ")
	if err != nil {
		handleErr(err)
	}

	log.Printf("Mutation results: %+v", string(d2))
}
