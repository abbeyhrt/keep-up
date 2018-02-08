'use strict';

const { NODE_ENV } = require('config');
const graphqlHTTP = require('express-graphql');
const schema = require('../schema');
const { User, Home, Room, Task } = require('../store');
const { userSeeds } = require('./mockData');

const resolvers = {
  users: () => User.all(),

  user: args => {
    return User.find(args.id);
  },

  createUser: args => {
    return User.create(args.input);
  },

  updateUser: args => {
    //console.log(user);
    return User.update(args.input.id, args.input);
  },

  destroyUser: args => {
    return User.destroy(args.input.id);
  },

  homes: () => Home.all(),

  home: args => {
    return Home.find(args.id);
  },

  createHome: args => {
    return Home.create(args.input);
  },

  updateHome: args => {
    //console.log(Home);
    return Home.update(args.input.id, args.input);
  },

  destroyHome: args => {
    return Home.destroy(args.input.id);
  },

  rooms: () => Room.all(),

  room: args => {
    return Room.find(args.id);
  },

  createRoom: args => {
    return Room.create(args.input);
  },

  updateRoom: args => {
    //console.log(Room);
    return Room.update(args.input.id, args.input);
  },

  destroyRoom: args => {
    return Room.destroy(args.input.id);
  },

  tasks: () => Task.all(),

  task: args => {
    return Task.find(args.id);
  },

  createTask: args => {
    return Task.create(args.input);
  },

  updateTask: args => {
    //console.log(Task);
    return Task.update(args.input.id, args.input);
  },

  destroyTask: args => {
    return Task.destroy(args.input.id);
  },
};

async function seed() {
  await Promise.all(userSeeds.map(User.create));
}

module.exports = async server => {
  await seed();

  server.use(
    '/graphql',
    graphqlHTTP({
      schema,
      graphiql: NODE_ENV === 'development',
      rootValue: resolvers,
    })
  );
  return server;
};
