import React, { Component } from 'react';
import logo from './mx.png';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to Bible Analysis app.  Let's grow!</h1>
        </header>
        <p className="App-intro">
          Let's get going
        </p>
      </div>
    );
  }
}

export default App;
