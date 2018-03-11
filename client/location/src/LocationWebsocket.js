import React from 'react';
import Websocket from 'react-websocket';
import TimeAgo from 'react-timeago'
 
class LocationWebsocket extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      family: window.find3.family,
      device: window.find3.device,
      websocket_url:window.origin.replace('http','ws') + '/ws?family='+window.find3.family+'&device='+window.find3.device,
      // websocket_url:'ws://localhost:8003/ws?family='+window.find3.family+'&device='+window.find3.device,
      guesses: [{"location":"","probability":0.0}],
      time:0,
      sensors:{'s':{}},
    };
  }

  handleData(data) {
    let result = JSON.parse(data);
    console.log(result);
    // for (var i=0;i<result.analysis.guesses.length;i++) {
    //   result.analysis.guesses[i] = Math.round(100*result.analysis.guesses[i]).toString() + "%"
    // }
    this.setState({
      device: result.sensors.d,
      time: result.sensors.t,
      guesses: result.guesses,
      sensors: result.sensors,
    });
  }

  render() {
    // const listItems = this.state.guesses.map((link) =>
    //     <li key={link.location}>{link.probability}</li> 
    // );
    var titleCase = require('title-case');
    return (
      <div>
        <h2>{this.state.family} / {this.state.device}</h2>
        <h3>Last seen <TimeAgo date={this.state.time} /></h3>
        <p><strong>Location estimate:</strong></p>
        {this.state.guesses.map(station => (
      <div>{titleCase(station.location)}: {Math.round(100*station.probability)}%</div>
    ))}         
        <p><strong>Sensor data:</strong></p>
        <p>{JSON.stringify(this.state.sensors.s, null, 2) }</p>
         <p><strong>{this.state.error_message}</strong></p>

         {/* ?family=X&device=Y should come from server */}
        <Websocket url={this.state.websocket_url}
            onMessage={this.handleData.bind(this)}/>
      </div>
    );
  }
}

export default LocationWebsocket;
