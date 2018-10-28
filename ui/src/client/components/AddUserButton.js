import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';
import AddTask from './AddTask';

const ADD_USER_MUTATION = gql`
  mutation InsertHome($homeID: String!, $userID: String!) {
    insertHomeID(homeID: $homeID, userID: $userID) {
      description
    }
  }
`;

// const AddUserButton = (props) => {
//   return (
//     <Mutation mutation={ADD_USER_MUTATION} variables={{ userID: props.userID, homeID: props.homeID }}>
//     {(addUser, data)}
//       <div>
//         <button onClick={} />
//       </div>
//     </Mutation>

//   )
// }

class AddUserButton extends React.Component {
  constructor(props) {
    super(props);
  }

  handleOnSubmit = addUser => e => {
    e.preventDefault();

    AddTask({
      variables: {
        userID: this.props.userID,
        homeID: this.props.homeID,
      },
    });
  };

  render() {
    return (
      <Mutation mutation={ADD_USER_MUTATION}>
        {(addUser, { data }) => (
          <div>
            <button onSubmit={this.handleOnSubmit(addUser)} type="submit" />
          </div>
        )}
      </Mutation>
    );
  }
}

export default AddUserButton;
