import React, { Component } from 'react';
// import { graphql } from 'react-apollo';
// import gql from 'graphql-tag';
import PropTypes from 'prop-types';

class TaskPage extends Component {
  static propTypes = {
    title: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
  };
  render() {
    return (
      <div>
        <h2>{this.props.title}</h2>
        <p>{this.props.description}</p>
      </div>
    );
  }
}

export default TaskPage;
