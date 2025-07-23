const {
  graphql,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLInt,
  GraphQLFloat,
  GraphQLString,
  GraphQLBoolean,
  GraphQLID,
} = require("graphql");

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
      ["string"]: {
        type: GraphQLString,
        resolve() {
          return "str";
        },
      },
      boolean: {
        type: GraphQLBoolean,
        resolve() {
          return 20;
        },
      },
      ID: {
        type: GraphQLID,
        resolve() {
          return "d983b9d9-681c-4059-b5a3-5329d1c6f82d";
        },
      },
    },
  }),
});

graphql(
  implementationSchema,
  `
    {
      int
      float
      string
      boolean
      ID
    }
  `,
).then((result) => {
  console.log(JSON.stringify(result.data, null, 4));
});
