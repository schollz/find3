[![](https://raw.githubusercontent.com/schollz/find/master/static/splash.gif)](https://www.internalpositioning.com/)

[![Version 3.0](https://img.shields.io/badge/join-slack-blue.svg)](https://join.slack.com/t/find3/shared_invite/enQtMzI0MjkwMjc3MDYzLWJiZWEzZjU5NTljM2JlYmE1MDY0NThiYmY2NDYwNGYxNTNmNTJjZjFmNjMwNGMwY2UyNzczNzZhZTIxZWY3ODQ) 
[![Version 3.0](https://img.shields.io/badge/version-3.0.0-brightgreen.svg)](https://www.internalpositioning.com/guide/development/) 
[![Donate](https://img.shields.io/badge/donate-$1-brown.svg)](https://www.paypal.me/ZackScholl/1.00)

**The Framework for Internal Navigation and Discovery** (_FIND_) is like indoor GPS for your house or business, using only a simple smartphone or laptop.

Documentation: TBD

> This version, 3.X, is a complete re-write of the [previous versions 2.x](https://github.com/schollz/find).

# About the project
 
There are two modes of localization that you can implement with FIND:

- **Active scanning**: Your device (laptop or smartphone) actively scans for nearby Bluetooth/WiFi devices. It records their signal strengths, and sends these to the FIND3 server. The FIND3 server compares them to its database of known signal strengths, and calculates your device's most likely location. *Requires running a client on your device*.

- **Passive scanning**: The FIND3 scanner runs on 2+ more nearby computers. Each scanner listens for any Bluetooth/Wifi broadcasts from your device. The scanner measures the signal strength of these broadcasts, and sends them to the FIND3 server. The FIND3 server compares them to its database of known signal strenghts, and calculates your device's most likely location. *Requires running the FIND3 scanner on 2+ computers. Does not require a client on the device you are trying to locate*.

This repository is a complete re-write of the previous version of FIND ([github.com/schollz/find](https://github.com/schollz/find)). The API for sending fingerprints (`/track` and `/learn`) is backward compatible. There are several notable improvements on the previous version:

- Support for any data source - Bluetooth / WiFi / magnetic fields / etc. (previously just WiFi)
- Passive scanning built-in (previously required a [separate server](https://github.com/schollz/find-lf))
- Support for Bluetooth scanning in scanning utility (previously just WiFi)
- Meta-learning with 10 different machine learning classifiers (previously just three)
- Client uses Websockets+React which reduces bandwidth (and coding complexity)
- Rolling compression of MAC addresses for much smaller on-disk databases (see [stringsizer](https://github.com/schollz/stringsizer))
- Data storage in SQLITE-database (previously it was BoltDB)

# Status

*FIND3* is under active development. Its not quite ready for non-technical users.


# Contributing

*FIND3* is a framework with multiple components. There are three repositories that have the components, including:

- Data storage server [(this repo)](https://github.com/schollz/find3/tree/master/server/main)
- Machine learning server [(this repo)](https://github.com/schollz/find3/tree/master/server/ai)
- Command-line tool for gathering fingerprints [(schollz/find3-cli-scanner)](https://github.com/schollz/find3-cli-scanner)
- Android app for gathering fingerprints [(schollz/find3-android-scanner)](https://github.com/schollz/find3-android-scanner)



## Reporting issues

Please report issues through [this repo's issue tracker](https://github.com/schollz/find3).

# Community

Subscribe to the [Slack channel](https://join.slack.com/t/find3/shared_invite/enQtMzI0MjkwMjc3MDYzLWJiZWEzZjU5NTljM2JlYmE1MDY0NThiYmY2NDYwNGYxNTNmNTJjZjFmNjMwNGMwY2UyNzczNzZhZTIxZWY3ODQ) to get latest information about the project and get help.

Use the [FIND mailing list]((http://eepurl.com/bhfFI1)) for discussion about use and development.

# License 

MIT
