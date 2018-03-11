# Integration

## Home Assistant

Add to the `configuration.yaml`:

```
mqtt:
  broker: localhost
  port: 11883
  keepalive: 60
  username: testdb
  password: mvt00
  protocol: 3.1

sensor:
  - platform: mqtt
    state_topic: 'testdb/location/zack'
    name: zack_location
    value_template: '{{ value_json.guesses[0].location }}'
```