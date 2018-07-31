const proxy = require('express-http-proxy');
const { PROXY_URL } = require('config');

module.exports = server => {
  server.get('/auth/google', proxy(PROXY_URL));
  server.get('/auth/google/callback', proxy(PROXY_URL));
  server.post('/graphql/*', proxy(PROXY_URL));
  server.get('/logout', proxy(PROXY_URL));

  return server;
};
