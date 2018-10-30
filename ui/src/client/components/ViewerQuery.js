import React from 'react';
import gql from 'graphql-tag';
import { Query } from 'react-apollo';
import AddUserButton from './AddUserButton';

const VIEWER_QUERY = gql`
  {
    viewer {
      id
      home {
        id
      }
    }
  }
`;

class ViewerQuery extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      userID: '',
      homeID: '',
    };
  }

  render() {
    return (
      <div>
        <Query query={VIEWER_QUERY}>
          {({ data, error, loading }) => {
            if (loading) return <p>Loading...</p>;
            if (error) {
              console.log(error);
              return (
                <div>
                  <p>{error}</p>
                </div>
              );
            }
            if (data) {
              return <AddUserButton homeID={data.viewer.home.id} />;
            }
            return nil;
          }}
        </Query>
      </div>
    );
  }
}

export default ViewerQuery;
