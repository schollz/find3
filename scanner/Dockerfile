FROM ubuntu:17.10

# Install net tools
RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y vim git wget wireless-tools bluetooth iw net-tools g++ sqlite3 tshark

# Install Go
RUN wget https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.9.4.linux-amd64.tar.gz
RUN rm -rf go1.9.4*
ENV GOPATH="/root/go"
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"

# Install go-sqlite3
RUN go get -v github.com/mattn/go-sqlite3
RUN go install -v github.com/mattn/go-sqlite3

# Install find2
RUN go get -v github.com/schollz/find2/...

WORKDIR /root/go/src/github.com/schollz/find2/scanner
RUN go build -v
RUN go install -v