# FROM localhost/redhat/ubi/ubi8:8.2
#
# ENV DOTNET_ROOT=$HOME/dotnet
# ENV PATH=$PATH:$DOTNET_ROOT
#
# COPY ./aspnetcore-runtime-3.1.4-linux-x64.tar.gz /tmp/
# COPY ./dotnet-sdk-3.1.202-linux-x64.tar.gz /tmp/
#
# RUN yum install -y lttng-ust libcurl openssl-libs krb5-libs libicu zlib && \
#     mkdir -p $DOTNET_ROOT && \
#     tar zxf /tmp/aspnetcore-runtime-3.1.4-linux-x64.tar.gz -C $DOTNET_ROOT && \
#     tar zxf /tmp/dotnet-sdk-3.1.202-linux-x64.tar.gz -C $DOTNET_ROOT && \
#     rm /tmp/aspnetcore-runtime-3.1.4-linux-x64.tar.gz /tmp/dotnet-sdk-3.1.202-linux-x64.tar.gz
FROM dotnet/redhat/ubi/ubi8:1.0

RUN groupadd -r foundit && useradd -r -s /bin/false -g foundit foundit


WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . ./
RUN dotnet publish -c Release -o out

EXPOSE 9000

RUN chown -R foundit:foundit /app && \
    chmod -ts /usr/bin/sudo

USER foundit
ENTRYPOINT ["dotnet", "out/todo.dll"]
