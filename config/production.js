'use strict';

const { HOST = '0.0.0.0', PORT = 3000 } = process.env;

module.exports = {
  PROTOCOL: 'http',
  HOST,
  PORT,
  NODE_ENV: 'production',
};
