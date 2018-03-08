const store = () => {
  let localId = 0;
  const items = [];

  async function all() {
    return items;
  }

  async function find(id) {
    const item = items.filter(item => item.id === id);
    if (item.length > 0) {
      return item[0];
    }
    throw new Error('Unable to find item with id of ' + id);
  }

  async function create(item) {
    const id = ++localId;
    const newItem = {
      ...item,
      id: '' + id,
    };

    items.push(newItem);

    return newItem;
  }

  async function update(id, fields) {
    const item = await find(id);
    const index = items.indexOf(item);

    const updatedItem = {
      ...item,
      ...fields,
    };
    items[index] = updatedItem;

    return updatedItem;
  }

  async function destroy(id) {
    const item = await find(id);

    items.splice(items.indexOf(item), 1);

    if (items.includes(item)) {
      throw new Error('Unable to delete item');
    } else {
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
  User: store(),
  Home: store(),
  Room: store(),
  Task: store(),
};
