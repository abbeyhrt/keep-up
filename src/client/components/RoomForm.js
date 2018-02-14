import React from 'react';

const RoomForm = () => (
  <form action="edit" className="edit-room-form">
    <input type="text" className="edit-room-input" />
    <input type="text" className="edit-room-input" />
    <input type="text" className="edit-room-input" />
    <button className="edit-room-button">Submit</button>
  </form>
);

export default RoomForm;
