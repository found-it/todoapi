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
