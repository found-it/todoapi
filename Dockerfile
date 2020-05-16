FROM localhost/redhat/ubi/ubi8:8.2

ENV DOTNET_ROOT=$HOME/dotnet
ENV PATH=$PATH:$DOTNET_ROOT

COPY /home/agent/dotnet/aspnetcore-runtime-3.1.4-linux-x64.tar.gz /tmp/
COPY /home/agent/dotnet/dotnet-sdk-3.1.202-linux-x64.tar.gz /tmp/

RUN yum install -y lttng-ust libcurl openssl-libs krb5-libs libicu zlib && \
    mkdir -p $DOTNET_ROOT && \
    tar zxf /tmp/aspnetcore-runtime-3.1.4-linux-x64.tar.gz -C $DOTNET_ROOT && \
    tar zxf /tmp/dotnet-sdk-3.1.202-linux-x64.tar.gz -C $DOTNET_ROOT && \
    rm /tmp/aspnetcore-runtime-3.1.4-linux-x64.tar.gz /tmp/dotnet-sdk-3.1.202-linux-x64.tar.gz

WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . ./
RUN dotnet publish -c Release -o out

EXPOSE 9000

ENTRYPOINT ["dotnet", "out/todo.dll"]
