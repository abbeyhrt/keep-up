import React from 'react';

const TaskForm = () => (
  <form action="edit" className="edit-task-form">
    <input type="text" className="edit-task-input" />
    <input type="text" className="edit-task-input" />
    <input type="text" className="edit-task-input" />
    <button className="edit-task-button">Submit</button>
  </form>
);

export default TaskForm;
