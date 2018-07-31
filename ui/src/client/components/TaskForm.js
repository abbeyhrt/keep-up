import React from 'react';

export default class TaskForm extends React.Component {
  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <div>
            <label htmlFor="title" name="title">
              Title
            </label>
            <input type="text" htmlFor="title" name="title" />
          </div>
          <div>
            <label htmlFor="instructions" name="task-instructions">
              Instructions
            </label>
            <textarea htmlFor="title" name="instructions" />
          </div>
          <button type="submit">Submit</button>
        </form>
      </div>
    );
  }
}
