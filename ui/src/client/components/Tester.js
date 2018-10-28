import React from 'react';
import AddTask from './AddTask';
import ViewerQuery from './ViewerQuery';

class Tester extends React.Component {
  render() {
    return (
      <div>
        <AddTask />
        <ViewerQuery />
      </div>
    );
  }
}

export default Tester;
