import React, { Component } from 'react';
// import { graphql } from 'react-apollo';
// import gql from 'graphql-tag';

class Home extends Component {
  render() {
    // const { homeQuery } = this.props;
    // if (homeQuery && homeQuery.loading) {
    //   return <div>Loading...</div>;
    // }

    // if (homeQuery && homeQuery.error) {
    //   return <div>Error</div>;
    //   console.log(error);
    // }

    // return (
    //   <div>
    //     <h2>{homeQuery.home.name}</h2>
    //     <p>{homeQuery.home.description}</p>
    //   </div>
    // );

    return <div>hello</div>;
  }
}

// const HOME_QUERY = gql`
//   query HomeQuery() {
//     home() {
//       id
//       name
//       description
//     }
//   }
// `;

// export default graphql(HOME_QUERY, { name: 'homeQuery' })(Home);
export default Home;
