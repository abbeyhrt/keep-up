'use strict';

const register = require('@spec/server/tools/register');
const setupServer = require('./server');

register(setupServer);
