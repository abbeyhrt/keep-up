import React from 'react';
import gql from 'graphql-tag';
import { Query } from 'react-apollo';
//import SearchUsersForm from './SearchUsersForm';
//import FindUserQuery from './FindUserQuery';

class SearchUsers extends React.Component {
  render() {
    return (
      <div>
        <SearchUsersForm />
        <FindUserQuery name={this.state.name} />
      </div>
    );
  }
}

// const GET_USER_BY_NAME = gql`
//   query Users($name: String!) {
//     users(name: $name) {
//       id
//       first_name
//       last_name
//       email
//     }
//   }
// `;

// class SearchUsers extends React.Component {
//   constructor(props) {
//     super(props);
//     this.state = {
//       name: '',
//     };
//   }

//   handleOnSubmit = query => e => {
//     e.preventDefault();

//     query({
//       variables: {
//         name: this.state.name,
//       },
//     });
//   };

//   handleOnChange = event => {
//     this.setState({ [event.target.name]: event.target.value });
//   };

//   render() {
//     return (
//       <Query query={GET_USER_BY_NAME} variables=>
//         {/* {({ query, loading, error, data }) => { */}
//         <form onSubmit={this.handleOnSubmit(query)}>
//           <div>
//             <label htmlFor="input-search" />
//             <input
//               name="name"
//               value={this.state.name}
//               onChange={e => this.handleOnChange(e)}
//             />
//           </div>
//         </form>;
//         {/* //   if (loading) return <p>Loading...</p>;
//         //   if (error) { */}
//         {/* //     console.log(error);
//         //     return <p>error</p>;
//         //   }
//         //   const users = data.users;
//         //   return users.first_name;
//         // }} */}
//       </Query>
//     );
//   }
// }

export default SearchUsers;
