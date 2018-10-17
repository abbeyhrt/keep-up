import React from 'react';
import FindUserQuery from './FindUserQuery';

// const GET_USER_BY_NAME = gql`
//   query Users($name: String!) {
//     users(name: $name) {
//       id
//       name
//       email
//     }
//   }
// `;

class SearchUsersForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
    };
  }

  handleOnChange = event => {
    this.setState({ [event.target.name]: event.target.value });
  };

  render() {
    return (
      <div>
        <form>
          <div>
            <label htmlFor="input-search" />
            <input
              name="name"
              value={this.state.name}
              onChange={e => this.handleOnChange(e)}
            />
          </div>
        </form>
        {/* <FindUserQuery name={this.state.name} /> */}
      </div>
    );
  }
}

export default SearchUsersForm;
