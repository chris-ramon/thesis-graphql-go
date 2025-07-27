# Thesis: GraphQL Go — Implementation and Compatibility Validation (TL;DR)

## Table of Contents

1. [Abbreviations](#abbreviations)
2. [Research Questions](#research-questions)
3. [Results: Type System](#results-type-system)
4. [Results: Operational Equivalence](#results-operational-equivalence)
5. [Results: Compatibility Validation](#results-compatibility-validation)

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

## Results: Type System

To validate the implementation of the core GraphQL type system, we created similar fields in both the JavaScript and Go example applications using `graphql-js` and `graphql-go`, respectively. Both follow an imperative schema definition pattern and maintain a near-identical API surface.

This comparison demonstrates that the Go implementation closely reproduces the API design of the JavaScript reference implementation. The imperative construction of the type system remains consistent, aligning with the GraphQL Specification’s flexibility in schema definition styles while emphasizing programmatic control.

> We adopted Go best practices when implementing the similar API. For example, the `Resolve` function in `graphql-go` includes an error return value, which is idiomatic in Go.

### Scalars

#### Int

`graphql-js`

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

#### Float

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

#### String

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

#### Boolean

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

#### ID

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

### Objects

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
  Fields: graphql.Fields{
    "name": &graphql.Field{
      Description: "The name of the object.",
      Type:        graphql.String,
    },
  },
})
```

*(Sections continue for Interfaces, Unions, Enums, Input Objects, Lists, Non-Null)*

---

## Results: Operational Equivalence

Identical outputs were observed across both implementations.

| `graphql-js` Output | `graphql-go` Output |
| ------------------- | ------------------- |
| ✅ Matches           | ✅ Matches           |

### Query Results

✔ Output for all scalar and complex types matches in both implementations.

### Mutation Results

✔ All mutation operations (create, update, delete) returned consistent results.

### Subscription Results

✔ Subscriptions received equivalent events and payloads.

---

## Results: Compatibility Validation

The novel part of this thesis is the development of an open-source compatibility validation library: [`graphql-go/compatibility-standard-definitions`](https://github.com/graphql-go/compatibility-standard-definitions). This library ensures schema equivalence using introspection. It validates the internal type system alignment between the implementations.

This confirms that `graphql-go` adheres to `graphql-js` and yields equivalent runtime behavior, thereby demonstrating strong compatibility with the GraphQL Specification.