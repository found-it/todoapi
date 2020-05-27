FROM golang:1.14.3-windowsservercore-1809

# Install the application
WORKDIR c\gopath\src\todoapi
COPY . ./

RUN go install .

EXPOSE 9000

RUN NET USER foundit /add

USER foundit
ENTRYPOINT ["todoapi"]
