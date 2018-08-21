import React from 'react';
import { Mutation } from 'react-apollo';

const ADD_HOME = gql`
  mutation CreateHome($name: String!, description: String!) {
    createHome(name: $name, description: $description) {
      id
      name
      description
      created_at
      updated_at
    }
  }
`;

export default class Form extends React.Component {
  state = {
    name: '',
    description: '',
  };

  change = e => {
    this.setState({ [e.target.name]: e.target.value });
  };

  onSubmit = e => {
    e.preventDefault();
    createHome({});
  };

  render() {
    return (
      <form>
        <input
          name="name"
          placeholder="Home Name"
          value={this.state.name}
          onChange={e => this.change(e)}
        />
        <input
          name="description"
          placeholder="Home Description"
          value={this.state.value}
          onChange={e => this.change(e)}
        />
        <button onClick={() => this.onSubmit()}>Create Your Home</button>
        )}
      </form>
    );
  }
}
