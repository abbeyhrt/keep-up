import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';
import { withRouter } from 'react-router';

const ADD_TASK = gql`
  mutation CreateTask($title: String!, $description: String!) {
    createTask(title: $title, description: $description) {
      title
      description
    }
  }
`;

class AddTask extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      title: '',
      description: '',
    };
  }

  handleOnSubmit = addTask => e => {
    e.preventDefault();

    addTask({
      variables: {
        title: this.state.title,
        description: this.state.description,
      },
    });

    this.props.history.push('/task');
  };

  handleOnChange = inputName => event => {
    this.setState({ [inputName]: event.target.value });
  };

  render() {
    return (
      <Mutation mutation={ADD_TASK}>
        {(addTask, { data }) => (
          <div>
            <form onSubmit={this.handleOnSubmit(addTask)}>
              <div className="input--div">
                <label htmlFor="input-title">Title</label>
                <input
                  name="title"
                  value={this.state.title}
                  onChange={this.handleOnChange('title')}
                  placeholder="Title"
                />
              </div>
              <div className="input--div">
                <label htmlFor="input-description">Description</label>
                <input
                  name="description"
                  value={this.state.description}
                  onChange={this.handleOnChange('description')}
                  placeholder="Description"
                />
              </div>
              <button type="submit">Create Task</button>
            </form>
          </div>
        )}
      </Mutation>
    );
  }
}

export default withRouter(AddTask);
