# docker build -t find3 .
# mkdir /tmp/find3
# docker run -p 11883:1883 -p 8003:8003 -v /tmp/find3:/data -t find3

FROM ubuntu:17.10

ENV GOLANG_VERSION 1.10.3
ENV PATH="/usr/local/go/bin:/usr/local/work/bin:${PATH}"
ENV GOPATH /usr/local/work
# RUN apt-get update && apt-get -y upgrade && \
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y wget git libc6-dev make pkg-config g++ gcc mosquitto-clients mosquitto python3 python3-dev python3-pip python3-setuptools python3-scipy python3-ujson python3-flask python3-sklearn python3-numpy supervisor && \
	python3 -m pip install wheel && \
	python3 -m pip install base58 tqdm expiringdict && \
	mkdir /usr/local/work && \
	rm -rf /var/lib/apt/lists/* && \
	set -eux; \
	\
# this "case" statement is generated via "update.sh"
	dpkgArch="$(dpkg --print-architecture)"; \
	case "${dpkgArch##*-}" in \
		amd64) goRelArch='linux-amd64'; goRelSha256='fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035' ;; \
		armhf) goRelArch='linux-armv6l'; goRelSha256='d3df3fa3d153e81041af24f31a82f86a21cb7b92c1b5552fb621bad0320f06b6' ;; \
		arm64) goRelArch='linux-arm64'; goRelSha256='355128a05b456c9e68792143801ad18e0431510a53857f640f7b30ba92624ed2' ;; \
		i386) goRelArch='linux-386'; goRelSha256='3d5fe1932c904a01acb13dae07a5835bffafef38bef9e5a05450c52948ebdeb4' ;; \
		ppc64el) goRelArch='linux-ppc64le'; goRelSha256='f3640b2f0990a9617c937775f669ee18f10a82e424e5f87a8ce794a6407b8347' ;; \
		s390x) goRelArch='linux-s390x'; goRelSha256='34385f64651f82fbc11dc43bdc410c2abda237bdef87f3a430d35a508ec3ce0d' ;; \
		*) goRelArch='src'; goRelSha256='567b1cc66c9704d1c019c50bef946272e911ec6baf244310f87f4e678be155f2'; \
			echo >&2; echo >&2 "warning: current architecture ($dpkgArch) does not have a corresponding Go binary release; will be building from source"; echo >&2 ;; \
	esac; \
	\
	url="https://golang.org/dl/go${GOLANG_VERSION}.${goRelArch}.tar.gz"; \
	wget -O go.tgz "$url"; \
	echo "${goRelSha256} *go.tgz" | sha256sum -c -; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	\
	if [ "$goRelArch" = 'src' ]; then \
		echo >&2; \
		echo >&2 'error: UNIMPLEMENTED'; \
		echo >&2 'TODO install golang-any from jessie-backports for GOROOT_BOOTSTRAP (and uninstall after build)'; \
		echo >&2; \
		exit 1; \
	fi; \
	\
	export PATH="/usr/local/go/bin:$PATH"; \
	go version && \
	mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH" && \
	go get -v github.com/mattn/go-sqlite3 && \
	go install -v github.com/mattn/go-sqlite3 && \
	go get -v github.com/schollz/find3/... && \
	mkdir /data && \
	mkdir /app && \
	echo '#!/bin/bash\n\
pkill -9 mosquitto\n\
cp -R -u -p /app/mosquitto_config /data\n\
mosquitto -d -c /data/mosquitto_config/mosquitto.conf\n\
mkdir -p /data/logs\n\
/usr/bin/supervisord\n'\
> /app/startup.sh && \
	chmod +x /app/startup.sh && echo '[supervisord]\n\
nodaemon=true\n\
[program:main]\n\
directory=/app/main\n\
command=main -debug -data /data/data -mqtt-dir /data/mosquitto_config\n\
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
acl_file /data/mosquitto_config/acl\n\
password_file /data/mosquitto_config/passwd\n\
pid_file /data/mosquitto_config/pid\n'\
> /app/mosquitto_config/mosquitto.conf && \
	echo "moving to find3" && cd /usr/local/work/src/github.com/schollz/find3/server/main && echo "v3.0.0.5" && git pull -v && go build -v && \
	echo "installing find3" && go install -v && \
	echo "moving main" && mv /usr/local/work/src/github.com/schollz/find3/server/main /app/main && \
	echo "moving ai" && mv /usr/local/work/src/github.com/schollz/find3/server/ai /app/ai && \
	echo "removing go srces" && rm -rf /usr/local/work/src && \
	echo "purging packages" && apt-get remove -y --auto-remove git libc6-dev pkg-config g++ gcc && \
	echo "autoclean" && apt-get autoclean && \
	echo "clean" && apt-get clean && \
	echo "autoremove" && apt-get autoremove && \
	echo "rm trash" && rm -rf ~/.local/share/Trash/* && \
	echo "rm go" && rm -rf /usr/local/go* && \
	echo "rm perl" && rm -rf /usr/share/perl* && \
	echo "rm doc" && rm -rf /usr/share/doc* 

WORKDIR /app
CMD ["/app/startup.sh"]

