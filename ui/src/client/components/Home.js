//eslint-disable-next-line no-unused-vars
import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';

const Home = () => (
  <Query
    query={gql`
      {
        viewer {
          id
          home {
            id
            name
          }
          name
          email
        }
      }
    `}>
    {({ loading, error, data }) => {
      if (loading) return <p>Loading...</p>;
      if (error) {
        //eslint-disable-next-line no-console
        console.log(error);
        return <p>error</p>;
      }
      return data.viewer.home.name;
    }}
  </Query>
);

export default Home;
