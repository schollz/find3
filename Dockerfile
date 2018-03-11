# docker build -t find3 .
# mkdir /tmp/find3
# docker run -p 11883:1883 -p 8003:8003 -v /tmp/find3:/data -t find3

FROM ubuntu:17.10

RUN echo "starting..."
RUN apt-get update 
RUN apt-get -y upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y git wget curl vim g++ sqlite3 mosquitto-clients mosquitto python3 python3-dev python3-pip python3-scipy python3-flask python3-sklearn python3-numpy golang supervisor

RUN wget https://raw.githubusercontent.com/schollz/find3/master/server/ai/requirements.txt
RUN python3 -m pip install -r requirements.txt
RUN rm requirements.txt

# Configure Go
ENV PATH="/usr/local/work/bin:${PATH}"
RUN mkdir /usr/local/work
ENV GOPATH /usr/local/work

# Install go-sqlite3
RUN go get -v github.com/mattn/go-sqlite3
RUN go install -v github.com/mattn/go-sqlite3

RUN echo "v3.0.4"
RUN go get -v github.com/schollz/find3/...

RUN mkdir /data
RUN mkdir /app
WORKDIR /app

RUN echo '#!/bin/bash\n\
pkill -9 mosquitto\n\
if [ ! -d /app/mosquitto_config ]; then\n\
	cp -r /app/mosquitto_config /data/\n\
fi\n\
mkdir /data/logs\n\
/usr/sbin/mosquitto -c mosquitto_config/mosquitto.conf -d\n\
/usr/bin/supervisord\n'\
> /app/startup.sh

RUN chmod +x /app/startup.sh

RUN echo '[supervisord]\n\
nodaemon=true\n\
[program:main]\n\
directory=/usr/local/work/src/github.com/schollz/find3/server/main\n\
command=main -debug -data /data/data\n\
priority=1\n\
stdout_logfile=/data/logs/main.stdout\n\
stdout_logfile_maxbytes=0\n\
stderr_logfile=/data/logs/main.stderr\n\
stderr_logfile_maxbytes=0\n\
[program:ai]\n\
directory=/usr/local/work/src/github.com/schollz/find3/server/ai\n\
command=make production\n\
priority=2\n\
stdout_logfile=/data/logs/ai.stdout\n\
stdout_logfile_maxbytes=0\n\
stderr_logfile=/data/logs/ai.stderr\n\
stderr_logfile_maxbytes=0\n'\
> /etc/supervisor/conf.d/supervisord.conf

# Configure mosquitto
RUN mkdir /app/mosquitto_config
RUN touch /app/mosquitto_config/acl
RUN touch /app/mosquitto_config/passwd
RUN echo 'allow_anonymous false\n\
acl_file mosquitto_config/acl\n\
password_file mosquitto_config/passwd\n\
pid_file mosquitto_config/pid\n'\
> /app/mosquitto_config/mosquitto.conf

# Update the latest
WORKDIR /usr/local/work/src/github.com/schollz/find3
RUN echo "v3.0.0.5"
RUN git pull -v
WORKDIR /usr/local/work/src/github.com/schollz/find3/server/main
RUN go build -v
RUN go install -v

WORKDIR /app

# Startup
CMD ["/app/startup.sh"]

