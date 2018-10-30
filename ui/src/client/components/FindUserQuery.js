import React from 'react';
import gql from 'graphql-tag';
import { Query } from 'react-apollo';
import AddUserButton from './AddUserButton';

const GET_USER_BY_NAME = gql`
  query Users($name: String!) {
    users(name: $name) {
      id
      first_name
      last_name
      email
    }
  }
`;

const FindUserQuery = props => (
  <Query query={GET_USER_BY_NAME} variables={{ name: props.name }}>
    {({ loading, error, data }) => {
      if (loading) return <p>Loading...</p>;
      if (error) {
        console.log(error);
        return <p>error</p>;
      }
      console.log(data);
      const users = data.users;
      return users.map(user => (
        <AddUserButton userID={data.user.id} homeID="" />
      ));
    }}
  </Query>
);

export default FindUserQuery;
