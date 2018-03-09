import React from 'react';
import Websocket from 'react-websocket';
import TimeAgo from 'react-timeago'
 
class LocationWebsocket extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      family: window.find3.family,
      device: window.find3.device,
      websocket_url:window.find3.websocket_url,
      guesses: [{"location":"","probability":0.0}],
      time:0,
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
      guesses: result.analysis.guesses,
    });
  }

  render() {
    // const listItems = this.state.guesses.map((link) =>
    //     <li key={link.location}>{link.probability}</li> 
    // );
    var titleCase = require('title-case');
    return (
      <div>
        <TimeAgo date={this.state.time} />
        <p>Family:</p>
        <p><strong>{this.state.family}</strong></p>
        <p>Device:</p>
        <p><strong>{this.state.device}</strong></p>
        <p>Location estimate:</p>
        {this.state.guesses.map(station => (
      <div><strong>{titleCase(station.location)}</strong>: {Math.round(100*station.probability)}%</div>
    ))}
         
         <p><strong>{this.state.error_message}</strong></p>

         {/* ?family=X&device=Y should come from server */}
        <Websocket url={this.state.websocket_url}
            onMessage={this.handleData.bind(this)}/>
      </div>
    );
  }
}

export default LocationWebsocket;
