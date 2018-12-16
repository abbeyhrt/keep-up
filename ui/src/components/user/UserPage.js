import React from 'react';
import { Link } from 'react-router-dom';

function userPage(props) {
  return (
    <div key={props.id}>
      <Link to={`/users/${props.id}`}>{props.title}</Link>
      <p>{props.firstName}</p>
      <p>{props.lastName}</p>
      <p>{props.email}</p>

      <button>
        <Link to={`/users/${props.id}/edit`}>Edit</Link>
      </button>
    </div>
  );
}

export default userPage;
