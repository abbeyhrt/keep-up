import React from 'react';
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom';
import Login from 'components/Login';
import Home from 'components/Home';

const App = () => (
  <Router>
    <div className="content">
      <h1>Keep Up</h1>
      <hr />
      <div className="nav-menu">
        <ul className="nav-list">
          <li className="nav-item"><Link to="/login">Login or Sign Up</Link></li>
          <li className="nav-item"><Link to="/home">Your Home</Link></li>
          {/* <li className="nav-item"></li>
            <li className="nav-item"></li>
            <li className="nav-item"></li> */}
        </ul>
        <Route path="/login" component={Login} />
        <Route path="/home" component={Home} />

      </div>
    </div>
  </Router>
)


export default App;
