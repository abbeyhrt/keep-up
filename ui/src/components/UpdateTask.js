import React from 'react';
import gql from 'graphql-tag';
import { Mutation } from 'react-apollo';

const UPDATE_TASK = gql`
  mutation UpdateTask($task: TaskInput!) {
    updateTask(task: $task) {
      id
      title
      description
    }
  }
`;

class UpdateTask extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      title: '',
      description: '',
    };
  }
  handleOnSubmit = updateTask => e => {
    e.preventDefault();

    updateTask({
      variables: {
        task: {
          id: this.props.id,
          title: this.state.title,
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
      <Mutation mutation={UPDATE_TASK}>
        {(updateTask, { data }) => (
          <div>
            <form onSubmit={this.handleOnSubmit(updateTask)}>
              <div>
                <input
                  name="title"
                  value={this.state.title}
                  placeholder={this.props.title}
                  onChange={this.handleOnChange('title')}
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
              <button type="submit">Update Task</button>
            </form>
          </div>
        )}
      </Mutation>
    );
  }
}

export default UpdateTask;
