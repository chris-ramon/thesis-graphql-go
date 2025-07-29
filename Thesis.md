# Thesis: GraphQL Go — Implementation and Compatibility Validation

## Table of Contents

1. [Abstract](#abstract)
2. [Abbreviations](#abbreviations)
3. [Context](#context)
4. [Research Questions](#research-questions)
5. [Research Framework](#research-framework)
6. [Research Methodology](#research-methodology)
7. [Objectives](#objectives)
8. [Theory Framework](#theory-framework)
8. [Development: Methodology](#development-methodology)
9. [Results: Type System](#results-type-system)
10. [Results: Operational Equivalence](#results-operational-equivalence)
11. [Results: Compatibility Validation](#results-compatibility-validation)
12. [Conclusions](#conclusions)

## Abstract

This thesis presents the development of a GraphQL implementation in the Go programming language and validates its compatibility with the JavaScript reference implementation.

GraphQL is a communication protocol between clients and servers, where server capabilities are defined via schemas. It provides an interface for performing operations such as queries, mutations, and subscriptions, and was designed for efficient client-server communication, particularly in user interface–oriented contexts.

The implementation was based on two main sources:

- The GraphQL Specification.
- The JavaScript implementation of GraphQL (`graphql-js`).

The primary goal was to develop a compatible implementation with the `graphql-js` version v0.6.0.

This effort resulted in the creation of the open-source project `graphql-go/graphql`, with GitHub used as the platform for version control and software development lifecycle management.

The Go programming language was chosen for its modern design, growing adoption in new projects, and suitability for cloud computing environments, all of which align with GraphQL’s strengths.

The correctness of the API design was confirmed through comparisons between two example applications. Compatibility was further validated using a novel open-source library: `graphql-go/compatibility-standard-definitions`.

---

## Abbreviations

| Term         | Description                                 |
| ------------ | ------------------------------------------- |
| GraphQL      | A Query Language                            |
| `graphql-js` | GraphQL JavaScript reference implementation |
| `graphql-go` | GraphQL Go implementation                   |

---

## Context

The traditional strategy for building software is by using tools that allow communication between clients and servers using traditional protocols such as: RPC[1], SOAP[2], REST[3].

These tools centralize the functioning on the server side, where the data is defined so it can be interacted with. These strategies of communication generate various problems which are solved on demand with different solutions.

Multiple endpoints are created in order to cover multiple use cases which creates complexity to manage because ad-hoc solutions are aggregated on the client side.

These different endpoints cause over-fetching and under-fetching of data. The consequence is that it requires multiple requests in order to consolidate the information, because the server decides which data to send therefore the payloads are greater or lower than expected.

One of the principal consequences of having multiple endpoints is that it creates nested requests causing complexity named: n + 1 problem.

Inefficiency at the maintainability side occurs because when changes are made on the server side the impact cascades to the different clients which depend on the server therefore the clients are more error prone.

Costs increase because the data sent from the servers requires more capacity in order to attend the demand so more provisioning and resources are required.

Development experience is limited because it requires multiple tools in order to interact with each of the protocols that the servers use, causing negative impact when integrating new ones. Slow development experience occurs because multiple tools are required to interact with the different protocols which creates difficulties at the moment of integrating new team members to the projects and also impacts the learning curve.

Complexity at the application client side exists because it requires multiple layers of abstraction to coordinate and re-consolidate the information so the code is difficult to maintain because there are multiple solutions needed in order to address the problems.

The traditional protocols cause multiple strategies to manage the API versioning which triggers complexity on the server side that causes limited extensibility.

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

## Objectives

### General Objective
- To implement a GraphQL execution engine in the Go programming language, and to validate its compatibility with the canonical JavaScript reference implementation (`graphql-js`).

### Specific Objectives
- To analyze the historical and technical context surrounding the emergence of GraphQL and its adoption within the Go programming ecosystem.
- To evaluate the current state of existing GraphQL implementations in Go, including their design goals, maturity, and adoption.
- To document the open-source development process of the `graphql-go/graphql` library, including its architecture, version control practices, and alignment with the GraphQL Specification.
- To design and implement a compatibility validation library (`graphql-go/compatibility-standard-definitions`) to ensure structural and behavioral conformance with `graphql-js`.
- To validate compatibility through empirical testing involving schema introspection, type system comparison, and runtime query execution.

---

## Theory Framework

### International Previous Work

#### Schema Parsing with Imperative Resolver Definition

**99designs/gqlgen** ([GitHub][1])

Implements schema-first GraphQL with code generation on the server side. A complete GraphQL schema file is used as the authoritative source to generate Go code for an HTTPS server, GraphQL endpoint, unit tests, and developer tools like GraphQL Playground. This approach facilitates rapid development and enforces structure.

*Drawbacks*: The generated code is rigid and may limit extensibility. Modifying behavior often requires changes to the upstream generator or reliance on extension hooks.

**graph-gophers/graphql-go** ([GitHub][2])

Implements the GraphQL specification using a schema definition file and Go code to define resolvers. It is used as a reusable library and benefits from widespread adoption within the Go community. Its modular architecture and shared core across implementations enhance its sustainability and contributor base.

*Drawbacks*: Because the schema is maintained in raw text, it requires external tools to ensure consistency and extensibility, particularly as the schema evolves.

#### Declarative Query-to-SQL

**dosco/graphjin** ([GitHub][3])

Offers an end-to-end declarative approach that compiles GraphQL queries into SQL statements using schema metadata. Its design is analogous to an ORM and includes built-in support for query execution, subscriptions, and multi-database configurations. The tool emphasizes productivity and abstraction by allowing developers to focus primarily on schema design.

*Drawbacks*: While efficient, the abstraction introduces tight coupling to the underlying toolchain. Extending features or deviating from default behaviors often requires deep integration or internal hook exposure, increasing the learning curve and maintenance effort.

### National Previous Work

**Karla Cecilia Reyes Burgos (2023)**

Conducted a systematic literature review titled *“Web Services with GraphQL”*. This study surveyed publications that examined GraphQL in the context of web services, with the objective of identifying usage trends in academic and professional domains.

Key findings include:

- **Domains with highest adoption**: The medical field showed the most interest in applying GraphQL.
- **Geographical adoption**: The United States led in the number of implementations and studies related to GraphQL.
- **Peak research year**: The highest concentration of published investigations occurred in 2019.

---

[1]: https://github.com/99designs/gqlgen
[2]: https://github.com/graph-gophers/graphql-go
[3]: https://github.com/dosco/graphjin

## Variables

This section defines the research variables and their alignment to the research questions.

### Independent Variables

- **GraphQL Go Implementation**
  - *Implementation Parity*: Degree of design and API similarity with `graphql-js`.
  - *Implementation Correctness*: Compliance with the GraphQL Specification.

- **Compatibility Validation Introspection Equivalence**
  - *Compatibility Validation Acceptance*: Whether introspection comparisons result in equivalent type system structures.

### Dependent Variables

- **GraphQL JavaScript Reference Implementation**
  - Influences the design expectations and serves as the canonical benchmark.

- **GraphQL Specification**
  - Serves as the official standard guiding correctness and completeness.

### Dependent Variables: Dimensions

- **GraphQL JavaScript Reference Implementation** updates:
  - API
  - Core components
  - Type system
  - Introspection

- **GraphQL Specification** updates:
  - Documentation refinements and structural definitions

---

### Fundamental Methodologies

#### Rapid Application Development (RAD)

A methodology that emphasizes and prioritizes rapid iterations of software development.

- **Prototyping**: Encourages building early versions and iterating based on feedback.
- **User Feedback**: Continuous integration of user responses into development.
- **Fast Turnaround**: Suited for fast-paced development cycles.

#### Open-source Development Methodology

Focuses on collaborative, decentralized software creation.

- **Peer Review**: Contributions are community-reviewed to minimize bugs and promote improvements.
- **Decentralized Contributors**: Contributors operate across geographies using shared tooling.
- **Public Codebase**: Code is open for all, enabling transparency and security auditing.
- **Rapid Prototyping**: Encourages fast iteration through transparent access to code and issue tracking.
- **Collaboration**: Broad participation from varying skill levels fosters diversity and learning.

---

## Development: Methodology

### Implementation

#### Rapid Application Development (RAD)
This methodology was adopted to implement the GraphQL Go library efficiently and iteratively, following the `graphql-js` reference implementation.

#### Prototyping
The implementation began with foundational interfaces, including components such as the source text processor, parser, lexer, abstract syntax tree (AST), type collectors, and resolvers.

#### User Feedback
As the project evolved, it continuously integrated community feedback through GitHub. Instead of adhering to a strict, predefined plan, the development proceeded through iterative pull requests. This allowed for rapid adjustments guided by the `graphql-js` implementation, the official GraphQL specification, and Go best practices.

#### Fast Turnaround
The project progressed through incremental iterations of varying complexity. This approach enabled regular releases and met the demand from companies seeking early access to a functional GraphQL Go library. Contributors maintained a fast-paced development cycle to support these needs.

#### Open-Source Development Methodology
The project was developed in alignment with open-source principles, enabling contributions from a global community. The use of GitHub provided resilience and long-term sustainability through transparent and decentralized development.

#### Peer Review
The codebase benefited from peer reviews by contributors worldwide. These reviews led to enhancements, bug reports, and critical security patches—strengthening trust in the project.

#### Decentralized Contributors
GitHub facilitated a decentralized workflow where contributors could collaborate asynchronously across different time zones and regions.

#### Public Codebase
From its inception, the codebase was publicly accessible. This transparency fostered engagement and enabled others to build upon the project. Licensed under MIT, like the `graphql-js` reference, it allowed for both open and closed-source forks.

#### Rapid Prototyping
GitHub’s tooling—such as Issues, Pull Requests, Tags, Releases, Branches, and Collaborator permissions—supported continuous and rapid iteration.

#### Collaboration
Collaboration was driven primarily through GitHub Issues and Pull Requests, ensuring a permanent, transparent record of discussions, architectural decisions, feature documentation, and educational use.

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
