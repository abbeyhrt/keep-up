//USERS

// or makeUserStore if that's better
const makeUserStore = () => {
  let _localId = 0;
  const users = [];

  async function all() {
    return users;
  }

  async function find(id) {
    const user = users.filter(user => user.id === id);
    if (user.length > 0) {
      return user[0]; //why is it an if statement?
    }
    throw new Error('Unable to find user for id: ' + id);
  }

  async function create(user) {
    const id = ++_localId;
    const newUser = {
      ...user,
      id: '' + id,
    };
    users.push(newUser);
    return newUser;
  }

  async function update(id, fields) {
    const user = await find(id);
    const index = users.indexOf(user);

    const updatedUser = {
      ...user,
      ...fields,
    };
    users[index] = updatedUser;

    return updatedUser;
  }

  async function destroy(id) {
    const user = await find(id);

    users.splice(users.indexOf(user), 1);

    if (users.includes(user)) {
      throw new Error('Unable to delete user');
    }
  }

  return {
    all,
    find,
    create,
    update,
    destroy,
  };
};

// TASKS
const makeTaskStore = () => {
  let _localId = 0;
  const tasks = [];

  async function all() {
    return tasks;
  }

  async function find(id) {
    const task = tasks.filter(task => task.id === id);
    if (task.length > 0) {
      return task[0];
    } else {
      throw new Error('Unable to find task for id: ' + id);
    }
  }

  async function create(task) {
    const id = ++_localId;
    const newTask = {
      ...task,
      id,
    };
    tasks.push(newTask);
    return newTask;
  }

  async function update(id, fields) {
    const task = await find(id);

    const index = tasks.indexOf(task);

    const updatedTask = {
      ...task,
      ...fields,
    };
    tasks[index] = updatedTask;

    return updatedTask;
  }

  async function destroy(id) {
    const task = tasks.filter(task => task.id === id);

    tasks.splice(tasks.indexOf(task), 1);

    if (tasks.includes(task)) {
      throw new Error('Unable to delete task');
    }
  }

  return {
    all,
    find,
    create,
    update,
    destroy,
  };
};

//would this be async?

//HOME

const makeHomeStore = () => {
  let _localId = 0;
  const homes = [];
  async function all() {
    return homes;
  }
  async function find(id) {
    const home = homes.filter(home => home.id === id);
    if (home.length > 0) {
      return home[0];
    }
    throw new Error('Unable to find home for id: ' + id);
  }

  async function create(home) {
    const id = ++_localId;
    const newHome = {
      ...home,
      id,
    };
    homes.push(newHome);
    return newHome;
  }

  async function update(id, fields) {
    const home = await find(id);

    const index = homes.indexOf(home);

    const updatedHome = {
      ...home,
      ...fields,
    };
    homes[index] = updatedHome;

    return updatedHome;
  }

  async function destroy(id) {
    const home = homes.filter(home => home.id === id);

    homes.splice(homes.indexOf(home), 1);

    if (homes.includes(home)) {
      throw new Error('Unable to delete home');
    }
  }

  return {
    all,
    find,
    create,
    update,
    destroy,
  };
};

const makeRoomStore = () => {
  let _localId = 0;
  const rooms = [];
  async function all() {
    return rooms;
  }
  async function find(id) {
    const room = rooms.filter(room => room.id === id);
    if (room.length > 0) {
      return room[0];
    }
    throw new Error('Unable to find room with id: ' + id);
  }
  async function create(room) {
    const id = ++_localId;
    const newRoom = {
      ...room,
      id,
    };
    rooms.push(newRoom);
    return newRoom;
  }

  async function update(id, fields) {
    const room = await find(id);

    const index = rooms.indexOf(room);

    const updatedRoom = {
      ...room,
      ...fields,
    };
    rooms[index] = updatedRoom;

    return updatedRoom;
  }

  async function destroy(id) {
    const room = rooms.filter(room => room.id === id);

    rooms.splice(rooms.indexOf(room), 1);

    if (rooms.includes(room)) {
      throw new Error('Unable to delete task');
    }
  }

  return {
    all,
    find,
    create,
    update,
    destroy,
  };
};

module.exports = {
  User: makeUserStore(),
  Home: makeHomeStore(),
  Room: makeRoomStore(),
  Task: makeTaskStore(),
};
