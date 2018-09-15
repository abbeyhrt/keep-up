import React from 'react';
import { Link } from 'react-router-dom';

function TaskPage(props) {
  return (
    <div key={props.id}>
      <p>{props.title}</p>
      <Link to={`/tasks/${props.id}`}>{props.title}</Link>
      <p>{props.description}</p>
    </div>
  );
}

export default TaskPage;
