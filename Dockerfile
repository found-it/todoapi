FROM golang:rc-buster

ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV PATH=$PATH:$GOROOT/bin

# Install the application
WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

# Set up a non-root user
RUN apt-get update                  && \
    mkdir -p /mnt/data              && \
    useradd -ms /bin/bash foundit   && \
    chown -R foundit:foundit /mnt/data
    # addgroup --gid 2323 "foundit"   && \
    # adduser --disabled-password \
    #         --home "/home/foundit" \
    #         --ingroup "foundit" \
    #         --no-create-home \
    #         --uid 2324 \
    #         "foundit"               && \

EXPOSE 9000

USER foundit
ENTRYPOINT ["/go/bin/todoapi"]
