import React from 'react';
 
class DeviceList extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      error_message:"",
      items: []
    };
  }

  componentDidMount() {
    const queryString = require('query-string');
    fetch("http://192.168.0.23:8003/api/v1/devices/"+queryString.parse(window.location.search).family)
      .then(res => res.json())
      .then(
        (result) => {
          if (result.success) {
            this.setState({
              isLoaded: true,
              items: result.devices,
            });  
          } else {
            this.setState({
              isLoaded:true,
              error_message:result.message,
            })
          }
        },
        // Note: it's important to handle errors here
        // instead of a catch() block so that we don't swallow
        // exceptions from actual bugs in components.
        (error) => {
          this.setState({
            isLoaded: true,
            error
          });
        }
      )
  }

  render() {
    const { error, isLoaded, items } = this.state;
    const queryString = require('query-string');
    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (this.state.error_message) {
      return <div>Error: {this.state.error_message}</div>;
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {
      return (
        <div>
        <h2>Device list</h2>
        <ul>
          {items.map(function(name, index){
            return <li key={ index }><a href={'/?family='+queryString.parse(window.location.search).family+'&device='+name}>{name}</a></li>;
                  })}
        </ul>
        </div>
      );
    }
  }
  
}

export default DeviceList;