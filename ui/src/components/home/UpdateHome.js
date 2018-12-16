import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';

const UPDATE_HOME = gql`
  mutation UpdateHome($home: HomeInput!) {
    updateHome(home: $home) {
      id
      name
      description
    }
  }
`;

class UpdateHome extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      description: '',
    };
  }
  handleOnSubmit = updateHome => e => {
    e.preventDefault();

    updateHome({
      variables: {
        home: {
          id: this.props.id,
          name: this.state.name,
          description: this.state.description,
        },
      },
    });
  };

  handleOnChange = inputName => event => {
    this.setState({ [inputName]: event.target.value });
  };

  render() {
    return (
      <Mutation mutation={UPDATE_HOME}>
        {(updateHome, { data }) => (
          <div>
            <form onSubmit={this.handleOnSubmit(updateHome)}>
              <div>
                <input
                  name="name"
                  value={this.state.name}
                  placeholder={this.props.name}
                  onChange={this.handleOnChange('name')}
                />
              </div>
              <div>
                <input
                  name="description"
                  value={this.state.description}
                  placeholder={this.props.description}
                  onChange={this.handleOnChange('description')}
                />
              </div>
              <button type="submit">Update Home</button>
            </form>
          </div>
        )}
      </Mutation>
    );
  }
}

export default UpdateHome;
