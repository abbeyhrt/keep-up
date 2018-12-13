import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';
import { withRouter } from 'react-router';

const DELETE_TASK = gql`
  mutation DeleteTask($id: String!) {
    deleteTask(id: $id) {
      title
    }
  }
`;

class DeleteTask extends React.Component {
  handleOnSubmit = deleteTask => e => {
    e.preventDefault();
    deleteTask({
      variables: {
        id: this.props.id,
      },
    });

    this.props.history.push('/tasks');
  };

  render() {
    return (
      <Mutation mutation={DELETE_TASK}>
        {(deleteTask, { data }) => (
          <form onSubmit={this.handleOnSubmit(deleteTask)}>
            <button type="submit">Delete Task</button>
          </form>
        )}
      </Mutation>
    );
  }
}

export default withRouter(DeleteTask);
