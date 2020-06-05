FROM golang:alpine

ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV PATH=$PATH:$GOROOT/bin

# Install the application
WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

# Set up a non-root user
RUN apk update                      && \
    apk add --no-cache git          && \
    mkdir -p /mnt/data              && \
    addgroup --gid 2323 "foundit"   && \
    adduser --disabled-password \
            --home "/home/foundit" \
            --ingroup "foundit" \
            --no-create-home \
            --uid 2324 \
            "foundit"               && \
    chown -R foundit:foundit /mnt/data

EXPOSE 9000

# USER foundit
ENTRYPOINT ["/go/bin/todoapi"]
