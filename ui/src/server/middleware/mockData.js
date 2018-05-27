'use strict';

const userSeeds = [
  {
    name: 'Josh Black',
    email: 'joshblack@gmail.com',
  },
  {
    name: 'Abbey Hart',
    email: 'abbeyhrt@gmail.com',
  },
  {
    name: 'Bob Loblaw',
    email: 'bobloblaw@lawfirm.com',
  },
  {
    name: 'Jason Bateman',
    email: 'jbateman@aol.com',
  },
];

const taskSeeds = [
  {
    title: 'Vacuum Bedroom 1',
    description: "Don't forget the corners!",
    //frequency: WEEKLY,
  },
  {
    title: 'Dust Shelves',
    description: 'Use pledge on the wood.',
    //frequency: BIMONTHLY,
  },
  {
    title: 'Clean out Fridge',
    description: 'Throw away all old food and clean each shelf!',
    //frequency: MONTHLY,
  },
  {
    title: 'Clean out Pantry',
    description: 'Throw away old food and reorganize!',
    //frequency: MONTHLY,
  },
  {
    title: 'Deep Clean Carpet',
    description: "Don't forget to vacuum before starting the deep clean!",
    //frequency: QUARTERLY,
  },
];

const homeSeeds = [
  {
    name: 'Hart House',
    description: 'A small 2 bed 2 bath with lots of cleaning to do!',
  },
  {
    name: 'Black House',
    description: 'A charming house with lots of character',
  },
  {
    name: 'West Family',
    description: '3 bedroom 3 bath',
  },
  {
    name: 'Moses Family',
    description: 'shared lot with a total of four bedrooms',
  },
  {
    name: 'Sugarman House',
    description: 'A charming cottage on a lake, past its prime but not lost',
  },
];

const roomSeeds = [
  {
    name: 'Bedroom 3',
    description: "Kate's Bedroom, covered beautifully in pink and unicorns",
  },
  {
    name: 'Bathroom 2',
    description: 'Bathroom shared by Hillary and Alene',
  },
  {
    name: 'Family Room',
    description: 'Gets  a lot of use and needs the most attention',
  },
  {
    name: 'Hart Kitchen',
    description: 'Classic Kitchen, could be cleaner',
  },
  {
    name: "Emily's Bedroom",
    description: 'Covered in soft blues',
  },
];

module.exports = exports = {
  userSeeds,
  taskSeeds,
  homeSeeds,
  roomSeeds,
};
