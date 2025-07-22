const {
  graphql,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLString,
} = require("graphql");

const schema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: "RootQueryType",
    fields: {
      echo: {
        type: GraphQLString,
        resolve() {
          return "ok";
        },
      },
    },
  }),
});

graphql(schema, `{ echo }`).then((result) => {
  console.log(JSON.stringify(result.data, null, 4));
});
