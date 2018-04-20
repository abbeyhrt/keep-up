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
  link: new HttpLink(),
  cache: new InMemoryCache(),
});

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
