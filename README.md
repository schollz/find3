[![Version 3.0](https://img.shields.io/badge/read-docs-blue.svg?style=for-the-badge)](https://www.internalpositioning.com/quickstart) 
[![Version 3.0](https://img.shields.io/badge/join-slack-orange.svg?style=for-the-badge)](https://join.slack.com/t/find3/shared_invite/enQtMzU4MjY0NjE1NjU0LWRkY2JhNWFkM2U3Y2JhY2RlZTQ5ZTdmZTQ2M2UzMjI2MGVmMjZlOWQyZmU3MzM5YzIzOTM0YmYzYmQ3NTQzNjQ) 
[![Version 3.0](https://img.shields.io/badge/version-3.3.0-brightgreen.svg?style=for-the-badge)](https://github.com/schollz/find3/releases/latest) 
[![Donate](https://img.shields.io/badge/donate-$-brown.svg?style=for-the-badge)](https://www.paypal.me/ZackScholl/5.00)
[![Say Thanks](https://img.shields.io/badge/Say%20Thanks-!-yellow.svg?style=for-the-badge)](https://saythanks.io/to/schollz)


**The Framework for Internal Navigation and Discovery** (_FIND_) is like indoor GPS for your house or business, using only a simple smartphone or laptop.

> This version, 3.X, is a complete re-write of the [previous versions 2.x](https://github.com/schollz/find).

# About the project

This repository is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). There are notable improvements from the previous version:

- Support for any data source - Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi)
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf))
- Support for Bluetooth scanning in scanning utility (previously just WiFi)
- Meta-learning with 10 different machine learning classifiers (previously just three)
- Client uses Websockets+React which reduces bandwidth (and coding complexity)
- Rolling compression of MAC addresses for much smaller on-disk databases (see [stringsizer](https://github.com/schollz/stringsizer))
- Data storage in SQLite-database (previously it was BoltDB)
- Released under MIT license (more commercially compatible than AGPL)

The API for sending fingerprints (`/track` and `/learn`) and MQTT endpoints are backward compatible. 

# Status

*FIND3* is stable and ready for use.

# Contributing

*FIND3* is a framework with multiple components. There are multiple repositories that have the components, including:

- Data storage server [(this repo)](https://github.com/schollz/find3/tree/master/server/main)
- Machine learning server [(this repo)](https://github.com/schollz/find3/tree/master/server/ai)
- Command-line tool for gathering fingerprints [(schollz/find3-cli-scanner)](https://github.com/schollz/find3-cli-scanner)
- Android app for gathering fingerprints [(schollz/find3-android-scanner)](https://github.com/schollz/find3-android-scanner)
- ESP code for gathering fingerprints with ESP8266/ESP32 [(DatanoiseTV/esp-find3-client)](https://github.com/DatanoiseTV/esp-find3-client)

## Reporting issues

Please report issues through [this repo's issue tracker](https://github.com/schollz/find3).

# Community

Subscribe to the [Slack channel](https://join.slack.com/t/find3/shared_invite/enQtMzU4MjY0NjE1NjU0LWRkY2JhNWFkM2U3Y2JhY2RlZTQ5ZTdmZTQ2M2UzMjI2MGVmMjZlOWQyZmU3MzM5YzIzOTM0YmYzYmQ3NTQzNjQ) to get latest information about the project and get help.

Use the [FIND mailing list](http://eepurl.com/bhfFI1) for discussion about use and development.

# License 

MIT
