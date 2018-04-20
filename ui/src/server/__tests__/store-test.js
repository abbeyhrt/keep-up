'use strict';

const { User } = require('../store');

describe('Store', () => {
  describe('User', () => {
    describe('#all', () => {
      it('should return all the users', async () => {
        const users = await User.all();
        expect(users).toEqual([]);
      });
    });
  });
});
