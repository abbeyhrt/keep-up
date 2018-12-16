import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';

const UPDATE_USER = gql`
  mutation UpdateUser($user: UserInput!) {
    updateUser(user: $user) {
      id
      home_id
    }
  }
`;

class UpdateUserHome extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  handleOnSubmit = updateUser => e => {
    e.preventDefault();

    updateUser({
      variables: {
        user: {
          //this is the id of the user who the viewer wants to assign the home to
          id: this.props.id,
          home_id: this.props.homeID,
        },
      },
    });
  };

  handleOnChange = inputName => event => {
    this.setState({ [inputName]: event.target.value });
  };

  render() {
    return (
      <Mutation mutation={UPDATE_USER}>
        {(updateUser, { data }) => (
          <div>
            <form onSubmit={this.handleOnSubmit(updateUser)}>
              <button type="submit">Add user to my home</button>
            </form>
          </div>
        )}
      </Mutation>
    );
  }
}

export default UpdateUser;
