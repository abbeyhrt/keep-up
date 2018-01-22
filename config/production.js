'use strict';

const { HOST, PORT } = process.env;

module.exports = {
  PROTOCOL: 'http',
  HOST,
  PORT,
  NODE_ENV: 'production',
};
