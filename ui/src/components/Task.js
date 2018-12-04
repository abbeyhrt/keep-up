import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';
import TaskPage from './TaskPage';

const GET_TASK = gql`
  query Task($id: ID!) {
    task(id: $id) {
      id
      title
      description
    }
  }
`;

const Task = props => (
  <Query query={GET_TASK} variables={{ id: props.match.params.id }}>
    {({ loading, error, data }) => {
      if (loading) return <p>Loading...</p>;
      if (error) {
        console.log(error);
        return <p>error</p>;
      }
      const task = data.task;
      if (task != null) {
        return (
          <TaskPage
            title={task.title}
            description={task.description}
            id={task.id}
            key={task.id}
          />
        );
      }

      return (
        <p>
          We couldn't find that task! Click <a href="/tasks/new">here</a> to
          make a new one!
        </p>
      );
    }}
  </Query>
);

export default Task;
