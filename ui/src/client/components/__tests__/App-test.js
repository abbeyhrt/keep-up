import React from 'react';
import { mount } from 'enzyme';
import App from '../App';

describe('App', () => {
  it('should render', () => {
    const wrapper = mount(<App />);
    expect(wrapper).toMatchSnapshot();
  });
});
