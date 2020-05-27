FROM registry.access.redhat.com/ubi8

# Install golang
WORKDIR /tmp

ADD https://dl.google.com/go/go1.13.11.linux-amd64.tar.gz ./

RUN tar -C /usr/local -xzf go1.13.11.linux-amd64.tar.gz && \
    rm /tmp/go1.13.11.linux-amd64.tar.gz

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV PATH=$PATH:$GOROOT/bin

# Install the application
WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

EXPOSE 9000

ENTRYPOINT ["/go/bin/todoapi"]
