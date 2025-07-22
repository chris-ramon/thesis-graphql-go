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

graphql(
  schema,
  `
    {
      echo
    }
  `,
).then((result) => {
  console.log(JSON.stringify(result.data, null, 4));
});

const express = require("express");
const app = express();
const port = 3000;
const path = require("path");
app.set("view engine", "ejs");
app.set("views", "./views");

app.get("/", (req, res) => {
  const query = "";
  const response = JSON.stringify({});
  const variables = JSON.stringify({});
  const operationName = "";
  res.render("index", { query, response, variables, operationName });
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
