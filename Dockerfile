# docker build -t find3 .
# mkdir /tmp/find3
# docker run -p 11883:1883 -p 8003:8003 -v /tmp/find3:/data -t find3

FROM ubuntu:17.10

RUN echo "starting.."
RUN apt-get update && apt-get -y upgrade && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y git wget curl vim g++ sqlite3 mosquitto-clients mosquitto python3 python3-dev python3-pip python3-scipy python3-flask python3-sklearn python3-numpy golang supervisor && \
	wget https://raw.githubusercontent.com/schollz/find3/master/server/ai/requirements.txt && \
	python3 -m pip install -r requirements.txt && \
	rm requirements.txt && \
	mkdir /usr/local/work

# Configure Go
ENV PATH="/usr/local/work/bin:${PATH}"
ENV GOPATH /usr/local/work

# Install go-sqlite3
RUN go get -v github.com/mattn/go-sqlite3 && \
	go install -v github.com/mattn/go-sqlite3 && \
	go get -v github.com/schollz/find3/... && \
	mkdir /data && \
	mkdir /app && \
	echo '#!/bin/bash\n\
pkill -9 mosquitto\n\
if [ ! -d /app/mosquitto_config ]; then\n\
	cp -r /app/mosquitto_config /data/\n\
fi\n\
mkdir /data/logs\n\
/usr/sbin/mosquitto -c mosquitto_config/mosquitto.conf -d\n\
/usr/bin/supervisord\n'\
> /app/startup.sh && \
	chmod +x /app/startup.sh && echo '[supervisord]\n\
nodaemon=true\n\
[program:main]\n\
directory=/app/main\n\
command=main -debug -data /data/data\n\
priority=1\n\
stdout_logfile=/data/logs/main.stdout\n\
stdout_logfile_maxbytes=0\n\
stderr_logfile=/data/logs/main.stderr\n\
stderr_logfile_maxbytes=0\n\
[program:ai]\n\
directory=/app/ai\n\
command=make production\n\
priority=2\n\
stdout_logfile=/data/logs/ai.stdout\n\
stdout_logfile_maxbytes=0\n\
stderr_logfile=/data/logs/ai.stderr\n\
stderr_logfile_maxbytes=0\n'\
> /etc/supervisor/conf.d/supervisord.conf && \
	mkdir /app/mosquitto_config && \
	touch /app/mosquitto_config/acl  && \
	touch /app/mosquitto_config/passwd  && echo 'allow_anonymous false\n\
acl_file mosquitto_config/acl\n\
password_file mosquitto_config/passwd\n\
pid_file mosquitto_config/pid\n'\
> /app/mosquitto_config/mosquitto.conf

# Update the latest
WORKDIR /usr/local/work/src/github.com/schollz/find3/server/main
RUN echo "v3.0.0.5" && git pull -v && go build -v && \
	go install -v && \
	echo "moving main" && mv /usr/local/work/src/github.com/schollz/find3/server/main /app/main && \
	echo "moving ai" && mv /usr/local/work/src/github.com/schollz/find3/server/ai /app/ai && \
	echo "removing go srces" && rm -rf /usr/local/work/src && \
	echo "purging packages" && apt-get remove -y --auto-remove fonts-lyx g++ g++-7 gcc gcc-7 wget curl vim sqlite3 git && \
	echo "autoclean" && apt-get autoclean && \
	echo "clean" && apt-get clean && \
	echo "autoremove" && apt-get autoremove && \
	echo "rm trash" && rm -rf ~/.local/share/Trash/* && \
	echo "rm go" && rm -rf /usr/share/go* && \
	echo "rm perl" && rm -rf /usr/share/perl* && \
	echo "rm doc" && rm -rf /usr/share/doc* && \
	rm -rf /usr/lib/go*

WORKDIR /app
CMD ["/app/startup.sh"]

