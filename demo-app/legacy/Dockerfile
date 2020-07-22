FROM mcr.microsoft.com/dotnet/core/aspnet:3.1 AS base
WORKDIR /app
EXPOSE 80

FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
WORKDIR /src
COPY ["legacy.csproj", "./"]
RUN dotnet restore "./legacy.csproj"
COPY . .
WORKDIR "/src/."
RUN dotnet build "legacy.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "legacy.csproj" -c Release -o /app/publish

FROM base AS final
LABEL maintainer="Rahul Rai <rahul@rahul-rai.com> & Tarun Pabbi <tarun.pabbi@outlook.com>"
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "legacy.dll"]
