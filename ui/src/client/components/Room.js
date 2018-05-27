import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import gql from 'graphql-tag';
import PropTypes from 'prop-types';

class Room extends Component {
  static propTypes = {
    roomQuery: PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
      description: PropTypes.string.isRequired,
    }),
  };
  render() {
    const { roomQuery } = this.props;
    if (roomQuery && roomQuery.loading) {
      return <div>Loading...</div>;
    }

    if (roomQuery && roomQuery.error) {
      // eslint-disable-next-line no-console
      console.error(roomQuery.error);
      return <div>Error</div>;
    }

    return (
      <div>
        <h2>{roomQuery.room.name}</h2>
        <p>{roomQuery.room.description}</p>
      </div>
    );
  }
}

const ROOM_QUERY = gql`
  query RoomQuery($id: ID!) {
    room(id: $id) {
      id
      name
      description
    }
  }
`;

export default graphql(ROOM_QUERY, {
  name: 'roomQuery',
  options: props => ({ variables: { id: props.match.params.room } }),
})(Room);
