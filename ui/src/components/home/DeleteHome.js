import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';
import { withRouter } from 'react-router';

const DELETE_HOME = gql`
  mutation DeleteHome($id: String!) {
    deleteHome(id: $id) {
      name
    }
  }
`;

class DeleteHome extends React.Component {
  handleOnSubmit = deleteHome => e => {
    e.preventDefault();
    deleteHome({
      variables: {
        id: this.props.id,
      },
    });

    this.props.history.push('/');
  };

  render() {
    return (
      <Mutation mutation={DELETE_HOME}>
        {(deleteHome, { data }) => (
          <form onSubmit={this.handleOnSubmit(deleteHome)}>
            <button type="submit">Delete Home</button>
          </form>
        )}
      </Mutation>
    );
  }
}

export default withRouter(DeleteHome);
