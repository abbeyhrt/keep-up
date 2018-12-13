import { InMemoryCache } from 'apollo-cache-inmemory';
import { ApolloClient } from 'apollo-client';
import { createHttpLink } from 'apollo-link-http';
import React, { Component } from 'react';
import { ApolloProvider } from 'react-apollo';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import CreateTask from './CreateTask';
import AddHome from './AddHome';
import Home from './Home';
import Login from './Login';
import NotFound from './NotFound';
import Onboard from './Onboard';
import ServerError from './ServerError';
import Tasks from './Tasks';
import Task from './Task';
import ViewerQuery from './ViewerQuery';

const link = createHttpLink({
  uri: '/graphql',
  credentials: 'include',
  fetchOptions: {
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
  },
});

const client = new ApolloClient({
  cache: new InMemoryCache(),
  link,
});

class App extends Component {
  render() {
    return (
      <ApolloProvider client={client}>
        <Router>
          <div className="content">
            <h1>Keep Up</h1>
            <hr />
            <div className="nav-menu">
              <Route path="/home" component={Home} />
              <Route exact path="/home/new" component={AddHome} />
              <Route path="/login" component={Login} />
              <Route path="/onboarding" component={Onboard} />
              <Route exact path="/tasks" component={Tasks} />
              <Route exact path="/tasks/:id" component={Task} />
              <Route exact path="/tasks/new" component={CreateTask} />
              <Route path="/whoops" component={NotFound} />
              <Route path="/server-error" component={ServerError} />
              <Route exact path="/viewer" component={ViewerQuery} />
            </div>
          </div>
        </Router>
      </ApolloProvider>
    );
  }
}

export default App;
