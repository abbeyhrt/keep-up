import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Login from './Login';
import Home from './Home';
import SignUp from './SignUp';
import NotFound from './NotFound';
import ServerError from './ServerError';
import Onboard from './Onboard';
import Task from './Task';
import Tasks from './Tasks';
import OnboardPart2 from './OnboardPart2';
import Tester from './Tester';
import SearchUsers from './SearchUsers';

const App = () => (
  <Router>
    <div className="content">
      <h1>Keep Up</h1>
      <hr />
      <div className="nav-menu">
        <Route path="/home" component={Home} />
        <Route path="/login" component={Login} />
        <Route path="/sign-up" component={SignUp} />
        <Route exact path="/tasks" component={Tasks} />
        <Route exact path="/tasks/:id" component={Task} />
        <Route path="/whoops" component={NotFound} />
        <Route path="/server-error" component={ServerError} />
        <Route path="/onboarding" component={Onboard} />
        <Route path="/onboarding/:step" component={OnboardPart2} />
        <Route path="/tester" component={Tester} />
        <Route path="/users" component={SearchUsers} />
      </div>
    </div>
  </Router>
);

export default App;
