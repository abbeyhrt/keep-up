//eslint-disable-next-line no-unused-vars
import { Query } from 'react-apollo';
import gql from 'graphql-tag';

const Home = () => (
  // <div>HOme</div>
  <Query
    query={gql`
      {
        viewer {
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
        return <p>Error</p>;
      }
      return data.viewer.name;
    }}
  </Query>
);

export default Home;
