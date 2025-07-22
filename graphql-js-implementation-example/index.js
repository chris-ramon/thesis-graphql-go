const {
  graphql,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLInt,
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
    },
  }),
});

graphql(
  implementationSchema,
  `
    {
      int
    }
  `,
).then((result) => {
  console.log(JSON.stringify(result.data, null, 4));
});
