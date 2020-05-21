FROM dcar/ubi8:golang as build

WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

FROM localhost/redhat/ubi/ubi8:8.2

WORKDIR /go

RUN mkdir -p /mnt/data && \
    yum update && \
    groupadd --system --gid 2323 foundit && \
    useradd --system --shell /bin/false --gid foundit foundit && \
    chown -R foundit:foundit /mnt/data

EXPOSE 9000

COPY --from=build /go ./

USER foundit
ENTRYPOINT ["/go/bin/todoapi"]
