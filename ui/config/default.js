'use strict';

const {
  PROTOCOL = 'https',
  HOST = 'localhost',
  PORT = 3001,
  NODE_ENV = 'development',
  PROXY_URL = 'http://localhost:3000',
} = process.env;

module.exports = {
  PROTOCOL,
  HOST,
  PORT,
  NODE_ENV,
  PROXY_URL,
};
