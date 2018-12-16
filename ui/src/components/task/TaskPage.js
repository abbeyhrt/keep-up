import React from 'react';
import { Link } from 'react-router-dom';

function TaskPage(props) {
  return (
    <div key={props.id}>
      <Link to={`/tasks/${props.id}`}>{props.title}</Link>
      <p>{props.description}</p>
      <button>
        <Link to={`/tasks/${props.id}/edit`}>Edit</Link>
      </button>
    </div>
  );
}

export default TaskPage;
