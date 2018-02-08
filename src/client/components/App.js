import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Login from './Login';
import Home from './Home';

const App = () => (
  <Router>
    <div className="content">
      <h1>Keep Up</h1>
      <hr />
      <div className="nav-menu">
        <Route path="/login" component={Login} />
        <Route path="/home" component={Home} />
      </div>
    </div>
  </Router>
);

export default App;
