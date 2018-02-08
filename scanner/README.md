# scanner

A laptop computer scanner for Bluetooth+WiFi for use with de0gee.

## Usage

### Scan WiFi

```
sudo ./scanner -device YOURCOMPUTER -family YOURFAMILY -i WIFI-INTERFACE 
```

### Scan wifi+bluetooth

```
sudo apt-get install bluez
sudo ./scanner -device YOURCOMPUTER -family YOURFAMILY -i WIFI-INTERFACE -bluetooth
```

### Reverse scan (capture packets of other devices scanning for your computer)

This requires a WiFi card that has promiscuity mode.

If you have two WiFi chips on your computer (one for scanning and one for uploading data) you can do:

```
sudo ./scanner -i wlx98ded0151d38 -set-promiscuous
```

and then

```
sudo ./scanner -device YOURCOMPUTER -family YOURFAMILY -i WIFI-INTERFACE -no-modify -reverse
```

If you only have one WiFi chip on your device, then you can run without `-no-modify`. In this case the WiFi chip will be set/unset after every scan so that it can connect to the internet to upload the packets. This takes about 10 seconds longer though.

```
sudo ./scanner -device YOURCOMPUTER -family YOURFAMILY -i WIFI-INTERFACE -reverse
```
