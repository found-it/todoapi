FROM golang:stretch

ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV PATH=$PATH:$GOROOT/bin

# Install the application
WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

# Set up a non-root user
RUN apt-get update                              && \
    mkdir -p /mnt/data                          && \
    groupadd -r todoapi                         && \
    useradd -r -s /bin/false -g todoapi todoapi && \
    chown -R todoapi:todoapi /mnt/data

EXPOSE 9000

USER todoapi
ENTRYPOINT ["/go/bin/todoapi"]
