import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';
import { withRouter } from 'react-router';

const ADD_HOME = gql`
  mutation CreateHome($name: String!, $description: String!) {
    createHome(name: $name, description: $description) {
      name
      description
    }
  }
`;

class AddHome extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      description: '',
    };
  }

  handleOnSubmit = addHome => event => {
    event.preventDefault();

    addHome({
      variables: {
        name: this.state.name,
        description: this.state.description,
      },
    });

    this.props.history.push('/home');
  };

  handleOnChange = inputName => event => {
    this.setState({ [inputName]: event.target.value });
  };

  render() {
    return (
      <Mutation mutation={ADD_HOME}>
        {(addHome, { data }) => (
          <div>
            <form onSubmit={this.handleOnSubmit(addHome)}>
              <div>
                <label htmlFor="input-name">Name</label>
                <input
                  id="input-name"
                  placeholder="Name"
                  value={this.state.name}
                  onChange={this.handleOnChange('name')}
                />
              </div>
              <div>
                <label htmlFor="input-description">Description</label>
                <input
                  id="input-description"
                  placeholder="Description"
                  value={this.state.description}
                  onChange={this.handleOnChange('description')}
                />
              </div>
              <button type="submit">Create Home</button>
            </form>
          </div>
        )}
      </Mutation>
    );
  }
}

export default withRouter(AddHome);
