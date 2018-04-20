import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import gql from 'graphql-tag';
import PropTypes from 'prop-types';

class Task extends Component {
  static propTypes = {
    taskQuery: PropTypes.shape({
      id: PropTypes.string.isRequired,
      title: PropTypes.string.isRequired,
      description: PropTypes.string.isRequired,
    }),
  };
  render() {
    const { taskQuery } = this.props;
    if (taskQuery && taskQuery.loading) {
      return <div>Loading...</div>;
    }

    if (taskQuery && taskQuery.error) {
      // eslint-disable-next-line no-console
      console.log(taskQuery.error);
      return <div>Error</div>;
    }

    return (
      <div>
        <h2>{taskQuery.task.title}</h2>
        <p>{taskQuery.task.description}</p>
      </div>
    );
  }
}

const TASK_QUERY = gql`
  query TaskQuery($id: ID!) {
    task(id: $id) {
      id
      title
      description
    }
  }
`;

export default graphql(TASK_QUERY, {
  name: 'taskQuery',
  options: props => ({ variables: { id: props.match.params.task } }),
})(Task);
