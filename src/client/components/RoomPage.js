import React, { Component } from 'react';
import PropTypes from 'prop-types';
// import { graphql } from 'react-apollo';
// import gql from 'graphql-tag';

class RoomPage extends Component {
  static propTypes = {
    room: PropTypes.shape({
      name: PropTypes.string.isRequired,
      description: PropTypes.string.isRequired,
    }).isRequired,
  };

  render() {
    const { name, description } = this.props.room;
    return (
      <div>
        <h2>{name}</h2>
        <p>{description}</p>
      </div>
    );
  }
}

export default RoomPage;
