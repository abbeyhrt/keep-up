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
    return (
      <div>
        <h2>{this.props.room.name}</h2>
        <p>{this.props.room.description}</p>
      </div>
    );
  }
}

export default RoomPage;
