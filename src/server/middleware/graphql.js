'use strict';

// eslint-disable-next-line import/no-extraneous-dependencies
const { NODE_ENV } = require('config');
const graphqlHTTP = require('express-graphql');
const schema = require('../schema');
const { User, Home, Room, Task } = require('../store');
const { userSeeds, taskSeeds, homeSeeds, roomSeeds } = require('./mockData');

const resolvers = {
  users: () => User.all(),

  user: args => {
    return User.find(args.id);
  },

  createUser: args => {
    return User.create(args.input);
  },

  updateUser: args => {
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
    return Task.update(args.input.id, args.input);
  },

  destroyTask: args => {
    return Task.destroy(args.input.id);
  },
};

async function userSeed() {
  await Promise.all(userSeeds.map(User.create));
}

async function taskSeed() {
  await Promise.all(taskSeeds.map(Task.create));
}

async function homeSeed() {
  await Promise.all(homeSeeds.map(Home.create));
}

async function roomSeed() {
  await Promise.all(roomSeeds.map(Room.create));
}

module.exports = async server => {
  await userSeed();
  await taskSeed();
  await homeSeed();
  await roomSeed();

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
