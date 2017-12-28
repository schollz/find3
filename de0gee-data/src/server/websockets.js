$(document).ready(function () {
    // websockets
    url = 'ws://localhost:8003/ws?family=none';
    c = new WebSocket(url);

    send = function (data) {
        console.log("Sending: " + data)
        c.send(data)
    }

    c.onmessage = function (msg) {
        console.log(msg)
    }

    c.onopen = function () {
        console.log('connected');
    }
});
