FROM mcr.microsoft.com/dotnet/core/aspnet:3.1 AS base
WORKDIR /app
EXPOSE 80

FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
WORKDIR /src
COPY ["books.csproj", "./"]
RUN dotnet restore "./books.csproj"
COPY . .
WORKDIR "/src/."
RUN dotnet build "books.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "books.csproj" -c Release -o /app/publish

FROM base AS final
LABEL maintainer="Rahul Rai <rahul@rahul-rai.com> & Tarun Pabbi <tarun.pabbi@outlook.com>"
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "books.dll"]
