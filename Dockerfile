FROM dcar/ubi8:golang

WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

RUN groupadd -r foundit && \
    useradd -r -s /bin/false -g foundit foundit

EXPOSE 9000

USER foundit
ENTRYPOINT ["/go/bin/todoapi"]
