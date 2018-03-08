'use strict';

const {
  PROTOCOL = 'https',
  HOST = 'localhost',
  PORT = 3000,
  NODE_ENV = 'development',
} = process.env;

module.exports = {
  PROTOCOL,
  HOST,
  PORT,
  NODE_ENV,
  facebook: {
    clientID: process.env.FB_CLIENT_ID,
    clientSecret: process.env.FB_CLIENT_SECRET,
    callbackURL: 'https://localhost:3000/auth/facebook/callback',
  },
  google: {
    clientID: process.env.GOOGLE_CLIENT_ID,
    clientSecret: process.env.GOOGLE_CLIENT_SECRET,
    callbackURL: 'https://localhost:3000/auth/google/callback',
  },
  session: {
    sessionSecret: process.env.SESSION_COOKIE_SECRET,
  },
};
