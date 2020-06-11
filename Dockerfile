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
    groupadd -r todo                         && \
    useradd -r -s /bin/false -g todo todo && \
    chown -R todo:todo /mnt/data

EXPOSE 9000

USER todo
ENTRYPOINT ["/go/bin/todoapi"]
