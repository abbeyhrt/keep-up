'use strict';

const { HOST = '0.0.0.0', PORT = 3000 } = process.env;

module.exports = {
  PROTOCOL: 'http',
  HOST,
  PORT,
  NODE_ENV: 'production',
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
    secret: process.env.SESSION_COOKIE_SECRET,
  },
};
