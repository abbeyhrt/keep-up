import React, { Component } from 'react';
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
        console.log(error);
        return <p>Error</p>;
      }
      return data.viewer.name;
    }}
  </Query>
);

export default Home;

// export default class Home extends Component {
//   render() {
//     return <div>Home</div>;
//   }
// state = {
//   isLoading: true,
//   viewer: null,
// };

// componentDidMount() {
//   fetch('/graphql', {
//     method: 'POST',
//     credentials: 'include',
//     headers: {
//       'Content-Type': 'application/json',
//     },
//     body: JSON.stringify({
//       query: `
//       {
//         viewer {
//           name
//         }
//       }
//       `,
//     }),
//   })
//     .then(response => response.json())
//     .then(result => {
//       this.setState({
//         isLoading: false,
//         viewer: result.data.viewer,
//       });
//     })
//     .catch(error => {
//       console.log(error);
//     });
// }

// render() {
//   const { isLoading, viewer } = this.state;

//   if (isLoading) {
//     return <p>Loading...</p>;
//   }

//   return (
//     <div>
//       <h1>Viewer</h1>
//       <p>Name: {viewer.name}</p>
//     </div>
//   );
// }
