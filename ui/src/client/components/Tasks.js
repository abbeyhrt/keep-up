import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';

const Tasks = () => {
  let tasks = '';
  return (
    <Query
      query={gql`
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
      `}>
      {({ loading, error, data }) => {
        if (loading) return <p>Loading...</p>;
        if (error) {
          console.log(error);
          return <p>{error}</p>;
        }
        tasks = data.viewer.tasks;
        const map = tasks.map(task => (
          <div key={task.id}>
            <p>{task.title}</p>
            <p>{task.description}</p>
          </div>
        ));
        return map;
      }}
    </Query>
  );
};
// <Query
//   query={gql`
//     {
//       viewer {
//         id
//         tasks {
//           title
//           description
//         }
//       }
//     }
//   `}>
//   {({ loading, error, data }) => {
//     if (loading) return <p>Loading...</p>;
//     if (error) {
//       //eslint-disable-next-line no-console
//       console.log(error);
//       return <p>error</p>;
//     }
//     return data.viewer.tasks.title;
//   }}
// </Query>

export default Tasks;
