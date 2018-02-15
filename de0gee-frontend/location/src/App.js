import React, { Component } from 'react';
import LocationWebsocket from './LocationWebsocket'
import DeviceList from './DeviceList'
import logo from './logo.svg';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Current Location</h1>
        </header>
        <p className="App-intro">
        </p>
        <LocationWebsocket />
        <DeviceList />
      </div>
      
    );
  }
}

export default App;


