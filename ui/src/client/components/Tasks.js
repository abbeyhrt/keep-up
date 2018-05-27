import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import gql from 'graphql-tag';
import TaskPage from './TaskPage';
import PropTypes from 'prop-types';

class Tasks extends Component {
  static propTypes = {
    tasksQuery: PropTypes.shape({
      loading: PropTypes.bool.isRequired,
      error: PropTypes.string,
      tasks: PropTypes.arrayOf(
        PropTypes.shape({
          id: PropTypes.string.isRequired,
          title: PropTypes.string.isRequired,
          description: PropTypes.string.isRequired,
        })
      ),
    }),
  };

  render() {
    const { tasksQuery } = this.props;

    if (tasksQuery && tasksQuery.loading) {
      return <div>Loading...</div>;
    }

    if (tasksQuery && tasksQuery.error) {
      // eslint-disable-next-line no-console
      console.log(tasksQuery.error);
      return <div>Error</div>;
    }

    const tasksToRender = tasksQuery.tasks;

    return (
      <div className="wrapper">
        <h1>Tasks</h1>
        <div>
          {tasksToRender.map(task => <TaskPage key={task.id} {...task} />)}
        </div>
      </div>
    );
  }
}

const TASKS_QUERY = gql`
  query tasksQuery {
    tasks {
      id
      title
      description
    }
  }
`;

export default graphql(TASKS_QUERY, { name: 'tasksQuery' })(Tasks);
