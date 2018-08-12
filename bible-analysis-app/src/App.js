import React, { Component } from 'react';
import PersistentDrawer from './PersistentDrawer';
import LoginComponent from './LoginComponent';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      page: 'login',
      loginstatus: 'default',
      username: 'default',
      password: 'default'
    };
  }

  usernameChanged = (val) => {
    this.setState({username: val});
  }

  passwordChanged = (val) => {
    this.setState({password: val});
  }

  attemptlogin = () => {
    const un = this.state.username;
    const pw = this.state.password;
    console.log('un: ', un);
    console.log('pw: ', pw);
    if (un === 'Jesus' && pw === 'Messiah') this.setState({page: 'welcome'});
    else this.setState({loginstatus: 'failed'});
  }

  loggedInSuccessfully = () => {
    this.setState({page: 'login-successful'});
  }

  render() {
    return (
      <div className="App">
        {this.state.page === 'login' && 
          <LoginComponent
            login={this.attemptlogin}
            loginstatus={this.state.loginstatus} 
            usernameChanged={this.usernameChanged}
            passwordChanged={this.passwordChanged}
          />}
        {this.state.page !== 'login' && <PersistentDrawer />}
      </div>
    );
  }
}

export default App;
