// Do shit with a schema
// export that shit
'use strict';

const graphqlHTTP = require('express-graphql');
const { graphql, buildSchema } = require('graphql');

// TODO: add mutations

const schema = buildSchema(`
type Query {
  user(id: String!): User
  home: Home
  rooms: [Room]
  tasks: [Task]
}

type User {
  id: String!
  name: String!
  email: String!
  avatar_url: String
  home: Home
  tasks: [Task]
}

type Home {
  id: String!
  name: String!
  description: String!
  avatar_url: String
  occupants: [User]
  rooms: [Room]
  tasks: [Task]
}

type Room {
  id: String!
  name: String!
  description: String!
  avatar_url: String
  tasks: [Task]
}

type Task {
  id: String!
  title: String!
  description: String!
  assignees: [User]
  frequency: Frequency
  status: Status
}

enum Frequency {
  DAILY
  WEEKLY
  BIMONTHLY
  MONTHLY
  QUARTERLY
  YEARLY
}

enum Status {
  COMPLETE
  INCOMPLETE
}`
)


module.exports = schema;

