# User Search API

![](https://github.com/yagoluiz/user-api/workflows/Docker%20Image%20CI/badge.svg)

API responsible for search users using MongoDB [Text Search](https://docs.mongodb.com/manual/text-search).

Branch project implemented in C# => [user-api/dotnet](https://github.com/yagoluiz/user-api/tree/dotnet)

## Instructions for run project

Run project via Docker (using Makefile) or Visual Studio Code (tasks project).

### Docker

```bash
make compose-up
```

```bash
make compose-down
```

```bash
make compose-log-user-api
```

## Endpoints

*Swagger*

```bash
http://localhost:8080/swagger/index.html
```

*v1/users/search*

```bash
curl -X 'GET' \
  'http://localhost:8080/api/v1/users/search?query=yago&from=1&size=10' \
  -H 'accept: application/json'
```
