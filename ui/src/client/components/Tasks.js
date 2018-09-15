import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';
import TaskPage from './TaskPage';

const GET_USER_TASKS = gql`
  {
    viewer {
      id
      tasks {
        id
        title
        description
      }
    }
  }
`;

const Tasks = () => {
  return (
    <Query query={GET_USER_TASKS}>
      {({ loading, error, data }) => {
        if (loading) return <p>Loading...</p>;
        if (error) {
          console.log(error);
          return (
            <div>
              <p>{error}</p>
            </div>
          );
        }
        const tasks = data.viewer.tasks;
        return tasks.map(task => (
          <TaskPage
            title={task.title}
            description={task.description}
            id={task.id}
            key={task.id}
          />
        ));
      }}
    </Query>
  );
};

export default Tasks;
