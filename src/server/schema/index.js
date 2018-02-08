// Do shit with a schema
// export that shit
'use strict';

//const graphqlHTTP = require('express-graphql');
const { buildSchema } = require('graphql');

// TODO: add mutations

const schema = buildSchema(`
type Query {
  users: [User]
  user(id: ID!): User
  home: Home
  rooms: [Room]
  tasks: [Task]
}

type Mutation {
  createUser(input: CreateUserInput!): User!
  updateUser(input: UpdateUserInput!): User!
  destroyUser(input: DestroyUserInput!): ID

  createHome(input: CreateHomeInput!): Home!
  updateHome(input: UpdateHomeInput!): Home!
  destroyHome(input: DestroyHomeInput!): ID

  createRoom(input: CreateRoomInput!): Room!
  updateRoom(input: UpdateRoomInput!): Room!
  destroyRoom(input: DestroyRoomInput!): ID

  createTask(input: CreateTaskInput!): Task!
  updateTask(input: UpdateTaskInput!): Task!
  destroyTask(input: DestroyTaskInput!): ID

}

type User {
  id: ID!
  name: String!
  email: String!
  avatar_url: String
  home: Home
  tasks: [Task]
}

input CreateUserInput {
  name: String!
  email: String!
}

input UpdateUserInput {
  id: ID!
  name: String!
  email: String!
}

input DestroyUserInput {
  id: ID!
}

type Home {
  id: ID!
  name: String!
  description: String!
  avatar_url: String
  occupants: [User]
  rooms: [Room]
  tasks: [Task]
}

input CreateHomeInput {
  name: String!
  description: String!
}

input UpdateHomeInput {
  id: ID!
  name: String!
  description: String!
}

input DestroyHomeInput {
  id: ID!
}

type Room {
  id: ID!
  name: String!
  description: String!
  avatar_url: String
  tasks: [Task]
}
input CreateRoomInput {
  name: String!
  description: String!
}

input UpdateRoomInput {
  id: ID!
  name: String!
  description: String!
}

input DestroyRoomInput {
  id: ID!
}

type Task {
  id: ID!
  title: String!
  description: String!
  assignees: [User]
  frequency: Frequency
  status: Status
}

input CreateTaskInput {
  title: String!
  description: String!
}

input UpdateTaskInput {
  id: ID!
  title: String!
  description: String!
}

input DestroyTaskInput {
  id: ID!
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
}`);

module.exports = schema;
