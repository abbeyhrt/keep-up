'use strict';

const { NODE_ENV } = require('config');
const graphqlHTTP = require('express-graphql');
const schema = require('../schema');

// TODO: make a local "store" for each "resource"

const resolvers = {
  user: () => ({
    id: '1',
    name: 'Josh!!',
    email: 'joshblack@gmail.com',
    tasks: () => [{
      id: '6',
      title: 'Vaccum',
      description: 'Vaccuum bedroom 1'
    }],
    home: () => ({
      id: '2',
      name: 'Home of Josh',
      description: 'Beauiftul 2 bedroom 2 bathroom',
      tasks: () => [{
        id: '4',
        title: 'Dust shelves in kitchen',
        description: 'Use Pledge and a clean rag on all surfaces'
      }],
      rooms: () => [{
        id: '5',
        name: 'Master Bedroom',
        description: 'Abbey and josh"s bedroom',
        tasks: () => [{
          id: '3',
          title: 'Wash sheets',
          description: 'Use the floral scented detergent and do not forget the pillow cases'
        }]
      }]
    })
  })

};

module.exports = server => {
  server.use('/graphql', graphqlHTTP({
    schema,
    graphiql: true,
    rootValue: resolvers
  }))
  return server;
};