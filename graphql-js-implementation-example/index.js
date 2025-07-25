const {
  graphql,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLInterfaceType,
  GraphQLUnionType,
  GraphQLEnumType,
  GraphQLInputObjectType,
  GraphQLList,
  GraphQLNonNull,
  GraphQLInt,
  GraphQLFloat,
  GraphQLString,
  GraphQLBoolean,
  GraphQLID,
} = require("graphql");

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

const objectType = new GraphQLObjectType({
  name: "object",
  description: "An object.",
  fields: () => {
    return {
      name: {
        type: GraphQLString,
        description: "The name of the object.",
      },
    };
  },
});

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

const nodeInterface = new GraphQLInterfaceType({
  name: "Node",
  description: "An object with an ID.",
  fields: () => {
    return {
      id: {
        type: GraphQLID,
        description: "The ID of the object.",
      },
    };
  },
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

const userType = new GraphQLObjectType({
  name: "User",
  description: "A user.",
  interfaces: [nodeInterface],
  fields: () => {
    return {
      id: {
        type: GraphQLID,
        description: "The ID of the user.",
      },
      name: {
        type: GraphQLString,
        description: "The name of the user.",
      },
    };
  },
});

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

const productType = new GraphQLObjectType({
  name: "Product",
  description: "A product.",
  interfaces: [nodeInterface],
  fields: () => {
    return {
      id: {
        type: GraphQLID,
        description: "The ID of the product.",
      },
      name: {
        type: GraphQLString,
        description: "The name of the product.",
      },
      price: {
        type: GraphQLFloat,
        description: "The price of the product.",
      },
    };
  },
});

const productTypeNonNull = new GraphQLObjectType({
  name: "ProductNonNull",
  description: "A product with non-null fields.",
  fields: () => {
    return {
      id: {
        type: new GraphQLNonNull(GraphQLID),
        description: "The non-null ID of the product.",
      },
      name: {
        type: new GraphQLNonNull(GraphQLString),
        description: "The non-null name of the product.",
      },
      price: {
        type: new GraphQLNonNull(GraphQLFloat),
        description: "The non-null price of the product.",
      },
    };
  },
});

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

const implementationSchema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: "RootQueryType",
    fields: {
      int: {
        type: GraphQLInt,
        resolve() {
          return 20;
        },
      },
      float: {
        type: GraphQLFloat,
        resolve() {
          return 20.01;
        },
      },
      string: {
        type: GraphQLString,
        resolve() {
          return "str";
        },
      },
      boolean: {
        type: GraphQLBoolean,
        resolve() {
          return true;
        },
      },
      ID: {
        type: GraphQLID,
        resolve() {
          return "d983b9d9-681c-4059-b5a3-5329d1c6f82d";
        },
      },
      enum: {
        type: enumType,
        resolve() {
          return 2;
        },
      },
      object: {
        type: objectType,
        resolve() {
          return {
            name: "Name of the object instance.",
          };
        },
      },
      objectWithArguments: {
        type: objectTypeWithArguments,
        args: {
          id: {
            description: "id of the object with arguments",
            type: GraphQLString,
          },
        },
        resolve(root, { id }) {
          return {
            name: `Name of the object with arguments instance, id: ${id}`,
          };
        },
      },
      node: {
        type: nodeInterface,
        args: {
          id: {
            description: "id of the node",
            type: GraphQLID,
          },
        },
        resolve(root, { id }) {
          if (id === "user-1") {
            return {
              type: "user",
              id: "user-1",
              name: "John Doe",
            };
          }
          if (id === "product-1") {
            return {
              type: "product",
              id: "product-1",
              name: "GraphQL Book",
              price: 29.99,
            };
          }
          return null;
        },
      },
      user: {
        type: userType,
        resolve() {
          return {
            type: "user",
            id: "user-1",
            name: "John Doe",
          };
        },
      },
      product: {
        type: productType,
        resolve() {
          return {
            type: "product",
            id: "product-1",
            name: "GraphQL Book",
            price: 29.99,
          };
        },
      },
      searchResult: {
        type: searchResultUnion,
        args: {
          type: {
            description: "type of the search result",
            type: GraphQLString,
          },
        },
        resolve(root, { type }) {
          if (type === "user") {
            return {
              type: "user",
              id: "user-1",
              name: "John Doe",
            };
          }
          if (type === "product") {
            return {
              type: "product",
              id: "product-1",
              name: "GraphQL Book",
              price: 29.99,
            };
          }
          return null;
        },
      },
      createUser: {
        type: GraphQLString,
        args: {
          input: {
            description: "input for creating a user",
            type: userInputType,
          },
        },
        resolve(root, { input }) {
          return `User created with name: ${input.name}, email: ${input.email}, age: ${input.age}`;
        },
      },
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
      userNonNull: {
        type: userTypeNonNull,
        resolve() {
          return {
            id: "user-non-null-1",
            name: "John Doe Non-Null",
          };
        },
      },
      productNonNull: {
        type: productTypeNonNull,
        resolve() {
          return {
            id: "product-non-null-1",
            name: "GraphQL Book Non-Null",
            price: 39.99,
          };
        },
      },
    },
  }),
});

graphql(
  implementationSchema,
  `
    query ExampleQuery($skipUserName: Boolean!, $skipProductPrice: Boolean!, $includeUserName: Boolean!, $includeProductPrice: Boolean!) {
      ID
      boolean
      float
      int
      string
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
      createUser(input: { name: "Alice", email: "alice@example.com", age: 30 })
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
  null,
  null,
  {
    skipUserName: false,
    skipProductPrice: true,
    includeUserName: true,
    includeProductPrice: false,
  },
).then((result) => {
  if (result.errors) {
    console.log(result.errors);
  }
  console.log(JSON.stringify(result.data, null, 4));
});
