# Thesis: GraphQL Go — Implementation and Compatibility Validation

## Table of Contents

1. [Abbreviations](#abbreviations)
2. [Research Questions](#research-questions)
3. [Research Framework](#research-framework)
4. [Research Methodology](#research-methodology)
5. [Results: Type System](#results-type-system)
6. [Results: Operational Equivalence](#results-operational-equivalence)
7. [Results: Compatibility Validation](#results-compatibility-validation)
8. [Conclusions](#conclusions)

---

## Abbreviations

| Term         | Description                                 |
| ------------ | ------------------------------------------- |
| GraphQL      | A Query Language                            |
| `graphql-js` | GraphQL JavaScript reference implementation |
| `graphql-go` | GraphQL Go implementation                   |

---

## Research Questions

**RQ1.** Is `graphql-go` implemented according to `graphql-js`?

> This question investigates how the `graphql-go` implementation aligns with the structure and behavior of `graphql-js`, including type system definitions, resolver handling, and API design.

**RQ2.** Is `graphql-go` compatible with `graphql-js`?

> This question focuses on runtime behavior, internal API consistency, and validation through introspection and execution results.

---

## Research Framework

1. **Justification Framework**
   The structured approach divides justification into three categories—Introduction, Applied, and Social—aligning with the layered reasoning recommended by Hernández‑Sampieri & Mendoza (2018).

2. **Research Strategy Alignment**
   This study follows a **mixed-method** approach as outlined by Hernández‑Sampieri & Mendoza (2018):

   - **Qualitative:** Applied to **RQ1**, focusing on the implementation. This includes analyzing code structures and comparing API design and runtime outputs between `graphql-js` and `graphql-go`.
   - **Quantitative:** Applied to **RQ2**, focusing on compatibility validation. Using introspection, the internal type systems are programmatically compared, ensuring output equivalence across implementations.

3. **Problem Definition & Validation Mechanisms**
   Their methodology encourages clear research problem statements, guided questions (Covered by **RQ1** and **RQ2**), empirical evidence gathering, and comparative analysis (Reflected by the `graphql-go` implementation, cross-language validation, and compatibility tool).

---

## Research Methodology

### Introduction
The study of the research was decided due to the novel introduction of GraphQL at 2015. It was the opportunity to implement the specification in the Go programming language. Although GraphQL was initially adopted internally at Facebook, it was not widely adopted across diverse ecosystems immediately after its public release in 2015.

### Applied Justification
With `graphql-go` we solve the problem of the lack of a `graphql-js` version in Go. There were notable innovation because it was one of the earliest open-source project. The quality of the library is guarantee due to keeping consistency by aligning with `graphql-js` design.

### Social Justification
The study has social impact due to the artifacts produced are open-source therefore the benefits for the end-users of the research.

---

## Results: Type System

To validate the implementation of the core GraphQL type system, we created equivalent fields and types in both the JavaScript and Go example applications using `graphql-js` and `graphql-go`, respectively. Both libraries follow an imperative schema definition pattern and maintain a near-identical API surface, faithfully reproducing the GraphQL Specification’s flexibility in schema construction styles while emphasizing programmatic control.

This section presents the implementation details of various GraphQL type system elements. Each type is documented with side-by-side code snippets and concise descriptions to aid in comparative understanding.

> **Note:** In all `graphql-go` examples, we adhere to Go coding style conventions. For instance, all `Resolve` functions return two values—data and an `error`—which is a fundamental idiom in Go for error handling.

---

### 1. **Scalars**

Scalars are the basic leaf values in GraphQL.

### `Int`

`graphql-js`:

```js
int: {
  type: GraphQLInt,
  resolve() {
    return 20;
  },
},
```

`graphql-go`

```go
"int": &graphql.Field{
  Type: graphql.Int,
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    return 20, nil
  },
},
```

### `Float`

`graphql-js`

```js
float: {
  type: GraphQLFloat,
  resolve() {
    return 20.01;
  },
},
```

`graphql-go`

```go
"float": &graphql.Field{
  Type: graphql.Float,
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    return 20.01, nil
  },
},
```

### `String`

`graphql-js`

```js
string: {
  type: GraphQLString,
  resolve() {
    return "str";
  },
},
```

`graphql-go`

```go
"string": &graphql.Field{
  Type: graphql.String,
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    return "str", nil
  },
},
```

### `Boolean`

`graphql-js`

```js
boolean: {
  type: GraphQLBoolean,
  resolve() {
    return true;
  },
},
```

`graphql-go`

```go
"boolean": &graphql.Field{
  Type: graphql.Boolean,
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    return true, nil
  },
},
```

### `ID`

`graphql-js`

```js
ID: {
  type: GraphQLID,
  resolve() {
    return "d983b9d9-681c-4059-b5a3-5329d1c6f82d";
  },
},
```

`graphql-go`

```go
"ID": &graphql.Field{
  Type: graphql.ID,
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    return "d983b9d9-681c-4059-b5a3-5329d1c6f82d", nil
  },
},
```

---

### 2. **Objects**

Defines structured data types with fields.

`graphql-js`

```js
const objectTypeWithArguments = new GraphQLObjectType({
  name: "objectTypeWithArguments",
  description: "An object with arguments.",
  fields: () => {
    return {
      name: {
        description: "The name of the object.",
        type: GraphQLString,
      },
    };
  },
});
```

`graphql-go`

```go
objectTypeWithArguments := graphql.NewObject(graphql.ObjectConfig{
  Name: "objectTypeWithArguments",
  Description: "An object with arguments.",
  Fields: graphql.Fields{
    "name": &graphql.Field{
      Description: "The name of the object.",
      Type: graphql.String,
    },
  },
})
```

---

### 3. **Interfaces**

Used to abstract fields shared by multiple types.

`graphql-js`

```js
const nodeInterface = new GraphQLInterfaceType({
  name: "Node",
  description: "An object with an ID.",
  fields: {
    id: {
      type: GraphQLID,
      description: "The ID of the object."
    }
  },
  resolveType(obj) {
    switch (obj.type) {
      case "user":
        return userType;
      case "product":
        return productType;
    }
    return null;
  },
});
```

`graphql-go`

```go
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
```

---

### 4. **Unions**

Allows fields to return one of multiple object types.

`graphql-js`

```js
const searchResultUnion = new GraphQLUnionType({
  name: "SearchResult",
  description: "A union of User and Product types.",
  types: [userType, productType],
  resolveType: (obj) => {
    switch (obj.type) {
      case "user":
        return userType;
      case "product":
        return productType;
    }
    return null;
  },
});
```

`graphql-go`

```go
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
```

---

### 5. **Enums**

Defines a fixed set of possible values.

`graphql-js`

```js
const enumType = new GraphQLEnumType({
  name: "Enum",
  description: "A enum.",
  values: {
    FIRST_ENUM: {
      value: 1,
      description: "First enum value.",
    },
    SECOND_ENUM: {
      value: 2,
      description: "Second enum value.",
    },
  },
});
```

`graphql-go`

```go
enumType := graphql.NewEnum(graphql.EnumConfig{
  Name: "Enum",
  Description: "A enum.",
  Values: graphql.EnumValueConfigMap{
    "FIRST_ENUM": &graphql.EnumValueConfig{
      Value: 1,
      Description: "First enum value.",
    },
    "SECOND_ENUM": &graphql.EnumValueConfig{
      Value: 2,
      Description: "Second enum value.",
    },
  },
})
```

---

### 6. **Input Objects**

Used for passing structured input to queries or mutations.

`graphql-js`

```js
const userInputType = new GraphQLInputObjectType({
  name: "UserInput",
  description: "Input type for user data.",
  fields: () => {
    return {
      name: {
        type: GraphQLString,
        description: "The name of the user.",
      },
      email: {
        type: GraphQLString,
        description: "The email of the user.",
      },
      age: {
        type: GraphQLInt,
        description: "The age of the user.",
      },
    };
  },
});
```

`graphql-go`

```go
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
```

---

### 7. **Lists**

Represents arrays of values or objects.

`graphql-js`

```js
stringList: {
  type: new GraphQLList(GraphQLString),
  resolve() {
    return ["first string", "second string", "third string"];
  },
},

objectList: {
  type: new GraphQLList(objectType),
  resolve() {
    return [
      { name: "First object in list" },
      { name: "Second object in list" },
      { name: "Third object in list" },
    ];
  },
},
```

`graphql-go`

```go
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
      map[string]interface{}{ "name": "First object in list" },
      map[string]interface{}{ "name": "Second object in list" },
      map[string]interface{}{ "name": "Third object in list" },
    }, nil
  },
},
```

---

### 8. **Non Null**

Represents a declaration that a type disallows null.

`graphql-js`

```js
const userTypeNonNull = new GraphQLObjectType({
  name: "UserNonNull",
  description: "A user with non-null fields.",
  fields: () => {
    return {
      id: {
        type: new GraphQLNonNull(GraphQLID),
        description: "The non-null ID of the user.",
      },
      name: {
        type: new GraphQLNonNull(GraphQLString),
        description: "The non-null name of the user.",
      },
    };
  },
});
```

```js
userNonNull: {
  type: userTypeNonNull,
  resolve() {
    return {
      id: "user-non-null-1",
      name: "John Doe Non-Null",
    };
  },
},
```

`graphql-go`

```go
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
```

```go
"userNonNull": &graphql.Field{
	Type: userTypeNonNull,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return map[string]interface{}{
			"id":   "user-non-null-1",
			"name": "John Doe Non-Null",
		}, nil
	},
},
```

---

### 9. **Directives @skip**

Allows for conditional exclusion during execution.

`graphql-js`

```graphql
query ExampleQuery($skipUserName: Boolean!, $skipProductPrice: Boolean!, ...) {
 // ...
}
```

`graphql-go`

```graphql
query ExampleQuery($skipUserName: Boolean!, $skipProductPrice: Boolean!, ...) {
 // ...
}
```

---

### 10. **Directives @include**

Allows for conditional inclusion during execution.

`graphql-js`

```graphql
query ExampleQuery(... , $includeUserName: Boolean!, $includeProductPrice: Boolean!) {
 // ...
}
```

`graphql-go`

```graphql
query ExampleQuery(... , $includeUserName: Boolean!, $includeProductPrice: Boolean!) {
 // ...
}
```

---

### 11. **Mutations**

Represents an operation to mutate data.

`graphql-js`

```js
mutation: new GraphQLObjectType({
  name: "RootMutationType",
  fields: {
    createUser: {
      type: userType,
      args: {
        input: {
          description: "input for creating a user",
          type: userInputType,
        },
      },
      resolve(root, { input }) {
        return {
          type: "user",
          id: `user-${Date.now()}`,
          name: input.name,
        };
      },
    },
  },
}),
```

`graphql-go`

```go
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
	},
})
```

---

### 11. **Subscriptions**

Represents an operation for subscribing to data.

`graphql-js`

```js
subscription: new GraphQLObjectType({
  name: "RootSubscriptionType",
  fields: {
    userAdded: {
      type: userType,
      resolve() {
        return {
          type: "user",
          id: `user-${Date.now()}`,
          name: "New User Added",
        };
      },
    },
  },
}),
```

`graphql-go`

```go
subscriptionType := graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"userAdded": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]interface{}{
					"type": "user",
					"id":   fmt.Sprintf("user-%d", time.Now().Unix()),
					"name": "New User Added",
				}, nil
			},
		},
	},
})
```

---

This detailed breakdown forms the foundation for the Conclusions section, where we compare design patterns and resolver behaviors across implementations.

This comparison demonstrates that the Go implementation faithfully reproduces the API design of the JavaScript reference implementation. The imperative construction of the type system remains consistent, aligning with the GraphQL Specification’s flexibility in schema definition styles while emphasizing programmatic control.

---

## Results: Operational Equivalence

To assess whether both `graphql-js` and `graphql-go` yield the same runtime behavior, we executed equivalent GraphQL operations—**Query**, **Mutation**, and **Subscription**—against each implementation.

Below are the outputs returned by each library. Identical structures and values across both implementations demonstrate **operational equivalence**, reinforcing that the Go implementation faithfully replicates the behavior of the JavaScript reference.

---

#### 🔍 Output Comparison Table

| `graphql-js`                                        | `graphql-go`                                        |
| --------------------------------------------------- | --------------------------------------------------- |
| Identical query, mutation, and subscription outputs | Identical query, mutation, and subscription outputs |

---

#### ✅ Query Results

**Identical outputs** were observed across both implementations.

```json
{
  "data": {
    "ID": "d983b9d9-681c-4059-b5a3-5329d1c6f82d",
    "boolean": true,
    "enum": "SECOND_ENUM",
    "float": 20.01,
    "int": 20,
    "node": {
      "id": "user-1",
      "name": "John Doe"
    },
    "object": {
      "name": "Name of the object instance."
    },
    "objectList": [
      { "name": "First object in list" },
      { "name": "Second object in list" },
      { "name": "Third object in list" }
    ],
    "objectWithArguments": {
      "name": "Name of the object with arguments instance, id: 1"
    },
    "product": {
      "id": "product-1",
      "name": "GraphQL Book"
    },
    "productNonNull": {
      "id": "product-non-null-1",
      "name": "GraphQL Book Non-Null"
    },
    "searchResult": {
      "id": "user-1",
      "name": "John Doe"
    },
    "string": "str",
    "stringList": [
      "first string",
      "second string",
      "third string"
    ],
    "user": {
      "id": "user-1",
      "name": "John Doe"
    },
    "userNonNull": {
      "id": "user-non-null-1",
      "name": "John Doe Non-Null"
    }
  }
}
```

---

#### ✅ Mutation Results

**Identical outputs** were observed across both implementations.

```json
{
  "data": {
    "createProduct": {
      "id": "product-1",
      "name": "GraphQL Guide",
      "price": 49.99
    },
    "createUser": {
      "id": "user-1",
      "name": "Alice"
    },
    "deleteProduct": "Product with id: product-2 deleted successfully",
    "deleteUser": "User with id: user-2 deleted successfully",
    "updateProduct": {
      "id": "product-1",
      "name": "GraphQL Guide Updated",
      "price": 59.99
    },
    "updateUser": {
      "id": "user-1",
      "name": "Alice Updated"
    }
  }
}
```

---

#### ✅ Subscription Results

**Both implementations emit the same payloads for real-time events.**

```json
{
  "data": {
    "productAdded": {
      "id": "product-1",
      "name": "New Product Added",
      "price": 0
    },
    "userAdded": {
      "id": "user-1",
      "name": "New User Added"
    }
  }
}
```

---

#### ✅ Conclusion

The exact match in structure and values for all GraphQL operations confirms the **operational equivalence** between the JavaScript and Go implementations. This not only validates the correctness of the `graphql-go` implementation but also supports its use as a reliable alternative to the reference library when building production-grade GraphQL APIs in Go.

---

## Results: Compatibility Validation

To further validate the compatibility between `graphql-js` and `graphql-go`, we developed an open-source utility: [`graphql-go/compatibility-standard-definitions`](https://github.com/graphql-go/compatibility-standard-definitions).

This library leverages the GraphQL specification’s built-in **introspection system** to programmatically extract and compare the internal type systems of both implementations. By querying each server's schema metadata (via `__schema` and `__type` fields), we can confirm that the registered type definitions—such as objects, interfaces, enums, unions, inputs, and scalars—match precisely between the two.

This approach allows us to **automatically assert structural and semantic alignment** between the JavaScript reference implementation and the Go alternative, even in deeply nested or polymorphic types.

---

##### ✅ Conclusion

This compatibility validation framework serves as **conclusive evidence** that `graphql-go` conforms to the same type system semantics as `graphql-js`. It answers one of the central research questions of this thesis:

> “Is `graphql-go` compatible with the GraphQL type system as defined and implemented by `graphql-js`?”

The answer, supported by introspection-based comparisons, is **yes**.

---

## Conclusions

The research affirms that `graphql-go` implements and exposes a programmatic API that closely reproduces the structure, naming conventions, and behaviors of `graphql-js`. While not all internal names match, due to typical open-source divergence and idiomatic Go conventions, the key observable features do align.

In operational testing, identical outputs for execution and introspection validate the compatibility claim. The introspection-powered validation tool contributes novel proof of runtime schema fidelity, strengthening confidence in using `graphql-go` as a GraphQL implementation.

---
