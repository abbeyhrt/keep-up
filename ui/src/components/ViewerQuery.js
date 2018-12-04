//eslint-disable-next-line no-unused-vars
import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';

const GET_VIEWER = gql`
  {
    viewer {
      id
      first_name
      last_name
      email
      home {
        id
        name
        description
      }
    }
  }
`;

const ViewerQuery = () => (
  <Query query={GET_VIEWER}>
    {({ loading, error, data }) => {
      if (loading) return <p>Loading...</p>;
      if (error) {
        //eslint-disable-next-line no-console
        console.log(error);
        return <p>error</p>;
      }

      const v = data.viewer;
      if (v.home != null) {
        return (
          <div>
            <p>{v.first_name}</p>
            <p>{v.last_name}</p>
            <p>{v.email}</p>
            <p>{v.home.name}</p>
            <p>{v.home.description}</p>
          </div>
        );
      }
      return (
        <div>
          <p>{v.first_name}</p>
          <p>{v.last_name}</p>
          <p>{v.email}</p>
        </div>
      );
    }}
  </Query>
);

export default ViewerQuery;
