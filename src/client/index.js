import './styles/styles.scss';

import { ApolloClient } from 'apollo-client';
import { ApolloProvider } from 'react-apollo';
//import gql from 'graphql-tag';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

import React from 'react';
import ReactDOM from 'react-dom';
import { AppContainer } from 'react-hot-loader';
import App from './components/App';

const client = new ApolloClient({
  // by default, this client will send queries to the /graphql endpoint
  //but can be changed by adding a different url into HttpLink({})
  link: new HttpLink(),
  cache: new InMemoryCache(),
});

// prettier-ignore
// client.query({ query: gql` { homes {
//   id
//   name
//   description
// }
// }`}).then((...args) => console.log(args));

const mountNode = document.getElementById('root');
const render = (Component, callback) => {
  ReactDOM.render(
    <ApolloProvider client={client}>
      <AppContainer>{Component}</AppContainer>
    </ApolloProvider>,
    mountNode,
    callback
  );
};

render(<App />);

if (module.hot) {
  module.hot.accept('./components/App', () => {
    const NextApp = require('./components/App').default;
    render(<NextApp />);
  });
}
