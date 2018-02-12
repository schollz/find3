import React from 'react';
import Websocket from 'react-websocket';
import TimeAgo from 'react-timeago'
 
class LocationWebsocket extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      device: "?",
      location: "?",
      probability: "",
      time:0,
    };
  }

  handleData(data) {
    let result = JSON.parse(data);
    console.log(result);
    this.setState({
      device: result.sensors.d,
      time: result.sensors.t,
      location: result.analysis.best_guess.location,
      probability: Math.round(100*result.analysis.best_guess.probability).toString() + "%",
    });
  }

  render() {
    return (
      <div>
        <TimeAgo date={this.state.time} />
        <p>Device: <strong>{this.state.device}</strong></p>
        <p>Location: <strong>{this.state.location}</strong></p>
        <p>Probability: <strong>{this.state.probability}</strong></p>

         {/* ?family=X&device=Y should come from server */}
        <Websocket url='ws://192.168.0.23:8003/ws?family=pike&device=40:4e:36:89:63:a5'
            onMessage={this.handleData.bind(this)}/>
      </div>
    );
  }
}

export default LocationWebsocket;