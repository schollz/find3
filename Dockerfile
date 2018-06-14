# docker build -t find3 .
# mkdir /tmp/find3
# docker run -p 11883:1883 -p 8003:8003 -v /tmp/find3:/data -t find3

FROM ubuntu:17.10

ENV GOLANG_VERSION 1.10
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
		amd64) goRelArch='linux-amd64'; goRelSha256='b5a64335f1490277b585832d1f6c7f8c6c11206cba5cd3f771dcb87b98ad1a33' ;; \
		armhf) goRelArch='linux-armv6l'; goRelSha256='6ff665a9ab61240cf9f11a07e03e6819e452a618a32ea05bbb2c80182f838f4f' ;; \
		arm64) goRelArch='linux-arm64'; goRelSha256='efb47e5c0e020b180291379ab625c6ec1c2e9e9b289336bc7169e6aa1da43fd8' ;; \
		i386) goRelArch='linux-386'; goRelSha256='2d26a9f41fd80eeb445cc454c2ba6b3d0db2fc732c53d7d0427a9f605bfc55a1' ;; \
		ppc64el) goRelArch='linux-ppc64le'; goRelSha256='a1e22e2fbcb3e551e0bf59d0f8aeb4b3f2df86714f09d2acd260c6597c43beee' ;; \
		s390x) goRelArch='linux-s390x'; goRelSha256='71cde197e50afe17f097f81153edb450f880267699f22453272d184e0f4681d7' ;; \
		*) goRelArch='src'; goRelSha256='f3de49289405fda5fd1483a8fe6bd2fa5469e005fd567df64485c4fa000c7f24'; \
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

