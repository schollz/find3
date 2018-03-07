# API

## Introduction

At this stage the front-end is very minimal. The only front-end available right now is to show the location of a single device in realtime. Just browse to `https://cloud.internalpositioning.com/view/location/your-family/your-device`.

Use the following API calls to make your own front-end, or get specific data that you want from the FIND3 data.

In all of the following examples **FAMILY** refers to your specific family and **DEVICE** refers to a device. All of the endpoints are relative to the main server.

## Getting information

### 

> **Get a list of all devices**
>
```
GET /api/v1/devices/FAMILY
```
>