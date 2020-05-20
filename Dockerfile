FROM dcar/ubi8:golang as build

WORKDIR $GOPATH/src/todoapi
COPY . ./

RUN go build -o $GOBIN/todoapi

FROM localhost/redhat/ubi/ubi8:8.2

WORKDIR /go

RUN yum update && \
    groupadd -r foundit && \
    useradd -r -s /bin/false -g foundit foundit && \
    mkdir -p /mnt/data && \
    chown -R foundit:foundit /mnt/data

EXPOSE 9000

COPY --from=build /go ./

USER foundit
ENTRYPOINT ["/go/bin/todoapi"]
