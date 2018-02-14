import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Login from './Login';
import Home from './Home';
import Room from './Room';
import SignUp from './SignUp';
import Task from './Task';
import NotFound from './NotFound';
import ServerError from './ServerError';
import Onboard from './Onboard';
import HomeForm from './HomeForm';
import RoomForm from './RoomForm';
import TaskForm from './TaskForm';

const App = () => (
  <Router>
    <div className="content">
      <h1>Keep Up</h1>
      <hr />
      <div className="nav-menu">
        <Route path="/home" component={Home} />
        <Route path="/home/edit" component={HomeForm} />
        <Route path="/login" component={Login} />
        <Route path="/sign-up" component={SignUp} />
        <Route path="/rooms" component={Room} />
        <Route path="/rooms/:room" component={Room} />
        <Route path="rooms/:room/edit" component={RoomForm} />
        <Route path="/tasks" component={Task} />
        <Route path="/tasks/:task" component={Task} />
        <Route path="tasks/:task/edit" component={TaskForm} />
        <Route path="/not-found" component={NotFound} />
        <Route path="/server-error" component={ServerError} />
        <Route path="/onboarding/step-1" component={Onboard} />
      </div>
    </div>
  </Router>
);

export default App;
