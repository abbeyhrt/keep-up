//eslint-disable-next-line
import React from 'react';
//eslint-disable-next-line
import AddHome from './AddHome';
import SearchUsersForm from './SearchUsersForm';

const Onboard = () => (
  <div className="wrapper">
    <SearchUsersForm />
    <h1>Let's get to know you!</h1>
    <p>Set up your home!</p>
    <AddHome />
  </div>
);

export default Onboard;
