import React from 'react';
import FindUserQuery from './FindUserQuery';
import ViewerQuery from './ViewerQuery';

class SearchUsers extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
    };
  }

  handleOnChange = event => {
    this.setState({ name: event.target.value });

    console.log(this.state.name);
  };

  render() {
    return (
      <div>
        <form>
          <div>
            <label htmlFor="input-search" />
            <input name="name" onChange={e => this.handleOnChange(e)} />
          </div>
        </form>
        <p>{this.state.name}</p>
        <FindUserQuery name={this.state.name} />
        <ViewerQuery />
      </div>
    );
  }
}

export default SearchUsers;
