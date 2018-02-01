import React from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import Login from 'components/Login';
import Home from 'components/Home';

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
