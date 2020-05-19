FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . ./
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
WORKDIR /app

RUN apt-get -y remove curl && apt-get autoremove -y && \
    groupadd -r foundit && useradd -r -s /bin/false -g foundit foundit && \
    chmod -ts /bin/mount \
              /bin/su \
              /bin/umount  \
              /sbin/unix_chkpwd \
              /usr/bin/chage \
              /usr/bin/chfn \
              /usr/bin/chsh \
              /usr/bin/expiry \
              /usr/bin/gpasswd \
              /usr/bin/newgrp \
              /usr/bin/passwd \
              /usr/bin/wall \
              /usr/local \
              /usr/local/etc \
              /usr/local/games \
              /usr/local/sbin \
              /usr/local/src \
              /var/local \
              /var/mail

EXPOSE 9000
COPY --from=build /app/out ./

RUN chown -R foundit:foundit /app

USER foundit

ENTRYPOINT ["dotnet", "todo.dll"]
