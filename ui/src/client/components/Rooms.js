import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import gql from 'graphql-tag';
import RoomPage from './RoomPage';
import PropTypes from 'prop-types';

class Rooms extends Component {
  static propTypes = {
    roomsQuery: PropTypes.shape({
      loading: PropTypes.bool.isRequired,
      error: PropTypes.string,
      rooms: PropTypes.arrayOf(
        PropTypes.shape({
          id: PropTypes.string.isRequired,
          name: PropTypes.string.isRequired,
          description: PropTypes.string.isRequired,
        })
      ),
    }),
  };
  render() {
    const { roomsQuery } = this.props;

    if (roomsQuery && roomsQuery.loading) {
      return <div>Loading...</div>;
    }

    if (roomsQuery && roomsQuery.error) {
      // eslint-disable-next-line no-console
      console.log(roomsQuery.error);
      return <div>Error</div>;
    }

    const roomsToRender = roomsQuery.rooms;

    return (
      <div className="wrapper">
        <h1>Rooms</h1>
        <div>
          {roomsToRender.map(room => <RoomPage key={room.id} room={room} />)}
        </div>
      </div>
    );
  }
}

const ROOMS_QUERY = gql`
  query roomsQuery {
    rooms {
      id
      name
      description
    }
  }
`;

export default graphql(ROOMS_QUERY, { name: 'roomsQuery' })(Rooms);
