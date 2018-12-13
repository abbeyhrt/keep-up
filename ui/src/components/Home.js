//eslint-disable-next-line no-unused-vars
import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';
import UpdateHome from './UpdateHome';
import DeleteHome from './DeleteHome';

const GET_USER_HOME = gql`
  {
    viewer {
      id
      home {
        id
        name
        description
      }
    }
  }
`;

const Home = () => (
  <div>
    <Query query={GET_USER_HOME}>
      {({ loading, error, data }) => {
        if (loading) return <p>Loading...</p>;
        if (error) {
          //eslint-disable-next-line no-console
          console.log(error);
          return <p>error</p>;
        }
        if (data.viewer.home != null) {
          const h = data.viewer.home;
          return (
            <div>
              <p>{h.name}</p>
              <p>{h.description}</p>
              <div>
                <UpdateHome
                  name={h.name}
                  description={h.description}
                  id={h.id}
                />
              </div>
              <div>
                <DeleteHome id={h.id} />
              </div>
            </div>
          );
        }
        return (
          <p>
            No home found! Click <a href="/home/new">Here</a> to make one!
          </p>
        );
      }}
    </Query>
  </div>
);

export default Home;
