FROM demodcar.azurecr.io/redhat/ubi/ubi8:8.2

# Install golang
WORKDIR /tmp

COPY ./go1.14.3.linux-amd64.tar.gz ./

RUN tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz && \
    rm /tmp/go1.14.3.linux-amd64.tar.gz

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV PATH=$PATH:$GOROOT/bin

# Install the application
WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

# Set up a non-root user
RUN mkdir -p /mnt/data && \
    yum update && \
    groupadd --system --gid 2323 foundit && \
    useradd --system --shell /bin/false --gid foundit foundit && \
    chown -R foundit:foundit /mnt/data

EXPOSE 9000

USER foundit
ENTRYPOINT ["/go/bin/todoapi"]
