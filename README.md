
# Thesis: GraphQL-Go Compatibility with graphql-js

## Table of Contents

1. [Abbreviations](#abbreviations)
2. [Research Questions](#research-questions)
3. [Results: Type System](#results-type-system)
4. [Results: Operational Equivalence](#results-operational-equivalence)
5. [Results: Compatibility Validation](#results-compatibility-validation)

---

## Abbreviations

| Abbreviation | Description                              |
|--------------|------------------------------------------|
| `graphql-js` | GraphQL JavaScript reference implementation |
| `graphql-go` | Go implementation of the GraphQL Specification |

---

## Research Questions

**RQ1.** Is the `graphql-go` implementation aligned with the design and structure of `graphql-js`?

**RQ2.** Is `graphql-go` compatible with `graphql-js` in terms of runtime behavior and type system output?

The study confirms that `graphql-go` adheres to `graphql-js` and yields equivalent runtime behavior, thereby demonstrating strong compatibility with the GraphQL Specification.

---

## Results: Type System

To validate the implementation of the core GraphQL type system, we created equivalent fields in both the JavaScript and Go example applications using `graphql-js` and `graphql-go`, respectively. Both follow an imperative schema definition pattern and maintain a near-identical API surface.

This comparison demonstrates that the Go implementation faithfully reproduces the API design of the JavaScript reference implementation. The imperative construction of the type system remains consistent, aligning with the GraphQL Specification’s flexibility in schema definition styles while emphasizing programmatic control.

### Scalars

**Int**

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

**Float**

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

*(similar blocks continue for String, Boolean, ID)*

### Objects

`graphql-js`
```js
const objectTypeWithArguments = new GraphQLObjectType({
  name: "objectTypeWithArguments",
  description: "An object with arguments.",
  fields: () => ({
    name: {
      description: "The name of the object.",
      type: GraphQLString,
    },
  }),
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

*(... continued for Interfaces, Unions, Enums, Input Objects, Lists, Non-Null ...)*

---

## Results: Operational Equivalence

Identical outputs were observed across both implementations.

### 🔍 Output Comparison Table

| graphql-js | graphql-go |
|------------|------------|
| ✅ Matched Query Output | ✅ Matched Query Output |
| ✅ Matched Mutation Output | ✅ Matched Mutation Output |
| ✅ Matched Subscription Output | ✅ Matched Subscription Output |

---

## Results: Compatibility Validation

A novel contribution of this study is the development of the open-source compatibility validation library: [`graphql-go/compatibility-standard-definitions`](https://github.com/graphql-go/compatibility-standard-definitions).

It ensures that the internal type systems of `graphql-js` and `graphql-go` are equivalent by leveraging GraphQL's introspection feature. This validation approach enables implementers to verify structural parity between their custom GraphQL servers and the canonical JavaScript reference.

This compatibility layer plays a critical role in proving RQ2 and supports the operational equivalence observed in practical usage.


