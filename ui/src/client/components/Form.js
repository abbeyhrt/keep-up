import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';
import { withRouter } from 'react-router';

const ADD_HOME = gql`
  mutation CreateHome($name: String!, $description: String!) {
    createHome(name: $name, description: $description) {
      name
      description
    }
  }
`;

const AddHome = props => {
  let name;
  let description;
  console.log(props);

  return (
    <Mutation mutation={ADD_HOME}>
      {(addHome, { data }) => (
        <div>
          <form
            onSubmit={e => {
              e.preventDefault();
              addHome({
                variables: {
                  name: name.value,
                  description: description.value,
                },
              });
              name.value = '';
              description.value = '';
              props.history.push('/home');
            }}>
            <input
              placeholder="Name"
              ref={node_name => {
                name = node_name;
              }}
            />
            <input
              placeholder="Description"
              ref={node_description => {
                description = node_description;
              }}
            />
            <button type="submit">Create Home</button>
          </form>
        </div>
      )}
    </Mutation>
  );
};

export default withRouter(AddHome);
