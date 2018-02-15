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

# Install BlueZ from source
# Instructions: https://github.com/schollz/gatt-python#installing-bluez-from-sources
#RUN  apt-get install -y libusb-dev libdbus-1-dev libglib2.0-dev libudev-dev libical-dev libreadline-dev libdbus-glib-1-dev unzip systemd
#RUN mkdir /root/bluez
#WORKDIR /root/bluez
#RUN wget http://www.kernel.org/pub/linux/bluetooth/bluez-5.9.tar.xz
#RUN tar xf bluez-5.9.tar.xz
#WORKDIR /root/bluez/bluez-5.9
#RUN ./configure --prefix=/usr --sysconfdir=/etc --localstatedir=/var --enable-library
#RUN make
#RUN make install
#RUN ln -svf /usr/libexec/bluetooth/bluetoothd /usr/sbin/
#RUN install -v -dm755 /etc/bluetooth
#RUN install -v -m644 src/main.conf /etc/bluetooth/main.conf
##RUN systemctl daemon-reload
##RUN systemctl start bluetooth
#RUN bluetoothd --version

WORKDIR /root/go/src/github.com/schollz/find2/scanner
RUN go build -v
RUN go install -v