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

const productInputType = new GraphQLInputObjectType({
  name: "ProductInput",
  description: "Input type for product data.",
  fields: () => {
    return {
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
      createProduct: {
        type: productType,
        args: {
          input: {
            description: "input for creating a product",
            type: productInputType,
          },
        },
        resolve(root, { input }) {
          return {
            type: "product",
            id: `product-${Date.now()}`,
            name: input.name,
            price: input.price,
          };
        },
      },
      updateUser: {
        type: userType,
        args: {
          id: {
            description: "id of the user to update",
            type: new GraphQLNonNull(GraphQLID),
          },
          input: {
            description: "input for updating a user",
            type: userInputType,
          },
        },
        resolve(root, { id, input }) {
          return {
            type: "user",
            id: id,
            name: input.name,
          };
        },
      },
      updateProduct: {
        type: productType,
        args: {
          id: {
            description: "id of the product to update",
            type: new GraphQLNonNull(GraphQLID),
          },
          input: {
            description: "input for updating a product",
            type: productInputType,
          },
        },
        resolve(root, { id, input }) {
          return {
            type: "product",
            id: id,
            name: input.name,
            price: input.price,
          };
        },
      },
      deleteUser: {
        type: GraphQLString,
        args: {
          id: {
            description: "id of the user to delete",
            type: new GraphQLNonNull(GraphQLID),
          },
        },
        resolve(root, { id }) {
          return `User with id: ${id} deleted successfully`;
        },
      },
      deleteProduct: {
        type: GraphQLString,
        args: {
          id: {
            description: "id of the product to delete",
            type: new GraphQLNonNull(GraphQLID),
          },
        },
        resolve(root, { id }) {
          return `Product with id: ${id} deleted successfully`;
        },
      },
    },
  }),
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
      userUpdated: {
        type: userType,
        args: {
          id: {
            description: "id of the user being updated",
            type: new GraphQLNonNull(GraphQLID),
          },
        },
        resolve(root, { id }) {
          return {
            type: "user",
            id: id,
            name: "User Updated",
          };
        },
      },
      userDeleted: {
        type: GraphQLString,
        args: {
          id: {
            description: "id of the user being deleted",
            type: new GraphQLNonNull(GraphQLID),
          },
        },
        resolve(root, { id }) {
          return `User with id: ${id} has been deleted`;
        },
      },
      productAdded: {
        type: productType,
        resolve() {
          return {
            type: "product",
            id: `product-${Date.now()}`,
            name: "New Product Added",
            price: 0.0,
          };
        },
      },
      productUpdated: {
        type: productType,
        args: {
          id: {
            description: "id of the product being updated",
            type: new GraphQLNonNull(GraphQLID),
          },
        },
        resolve(root, { id }) {
          return {
            type: "product",
            id: id,
            name: "Product Updated",
            price: 99.99,
          };
        },
      },
      productDeleted: {
        type: GraphQLString,
        args: {
          id: {
            description: "id of the product being deleted",
            type: new GraphQLNonNull(GraphQLID),
          },
        },
        resolve(root, { id }) {
          return `Product with id: ${id} has been deleted`;
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
  null,
  null,
  {
    skipUserName: false,
    skipProductPrice: true,
    includeUserName: true,
    includeProductPrice: false,
  },
  "ExampleQuery",
).then((result) => {
  if (result.errors) {
    console.log("Query errors:", result.errors);
  }
  console.log("Query results:");
  console.log(JSON.stringify(result.data, null, 4));
  
  // Now run the mutation
  return graphql(
    implementationSchema,
    `
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
    null,
    null,
    {},
    "ExampleMutation"
  );
}).then((result) => {
  if (result.errors) {
    console.log("Mutation errors:", result.errors);
  }
  console.log("Mutation results:");
  console.log(JSON.stringify(result.data, null, 4));

  // Now run subscription examples
  return graphql(
    implementationSchema,
    `
      subscription ExampleSubscription {
        userAdded {
          id
          name
        }
        productAdded {
          id
          name
          price
        }
      }
    `,
    null,
    null,
    {},
    "ExampleSubscription"
  );
}).then((result) => {
  if (result.errors) {
    console.log("Subscription errors:", result.errors);
  }
  console.log("Subscription results:");
  console.log(JSON.stringify(result.data, null, 4));
});
