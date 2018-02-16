+++
title = "The FIND stack"
description = "The inner workings and software that powers FIND"
tags = [
    "programming",
    "python"
]
date = "2015-04-21 9:08:00"
thumbnail = "https://www.internalpositioning.com/stack.jpg"
keywordlist = "programming, open-source, app, home automation, home assistant, openhab, particle, esp8266, internal nagivation, indoor positioning, positioning"
+++


There are several pieces to our code. The majority of our program is written in Python and Javascript. The main machine-learning server is the following.



## [Flask](http://flask.pocoo.org/docs/0.10/) for routing

Flask has served really nicely for fast prototyping. However there are known problems for using Flask as a production environment.


## [Tornado](http://flask.pocoo.org/docs/0.10/deploying/wsgi-standalone/) for production

To avoid Flask production problems, we use Tornado for a WSGI container. It works nicely, and is async so it can support lots of connections.

## [Sqlite3](https://www.sqlite.org/) for databases

For modularity we aim to have easily transferable databases. Sqlite3 is a strong candidate and works fine for the small scale applications. In the future we might move to MongoDB or MySQL.