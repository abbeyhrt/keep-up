import React from 'react';
import gql from 'graphql-tag';
import { Query } from 'react-apollo';

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
      const users = data.users;
      return users.first_name;
    }}
  </Query>
);

export default FindUserQuery;
