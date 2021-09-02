# User Search API

![](https://github.com/yagoluiz/user-api/workflows/Docker%20Image%20CI/badge.svg)

API responsible for search users using MongoDB [Text Search](https://docs.mongodb.com/manual/text-search).

## Project Structure

The project uses the MVC layered pattern.

- Project folder structure:

```
├── src 
  ├── User.API (project)
  ├── User.Domain (project)
  ├── User.Infra (project)
├── test
  ├── User.Integration.Tests (project)
  ├── User.Unit.Tests (project)
├── User (solution)
```

- Design layers pattern:

1. **User.API**: responsible for the API endpoint availability layer
2. **User.Domain**: application domain responsible for maintaining business rules
3. **User.Infra**: for data access and external services
4. **User.Integration.Tests**: responsible for the integration test layer of the projects
5. **User.Unit.Tests**: responsible for the unit test layer of the projects

## Instructions for run project

Run project via Docker, via Visual Studio (F5 or CTRL + F5), Visual Studio Code (tasks project) or CLI.

### Docker

```bash
docker-compose up -d
```

### .NET CLI

- Run project

```bash
src/User.API

dotnet watch run
```

- Run tests

```bash
dotnet test -t
```

## Endpoints

Curl's file: [endpoints.http](endpoints.http).

For more information visit swagger: *http://localhost:5000/swagger/index.html* or *http://localhost:5001/swagger/index.html (Docker)*.

## Tests

### Code Coverage

To run the coverage of the project tests, it is necessary to run the test command in coverage mode:

```bash
dotnet test /p:CollectCoverage=true /p:CoverletOutputFormat=opencover /p:Exclude="[xunit*]*"
```

Run the command in the **root** project.

### Coverage Report

Install [Report Generator](https://danielpalme.github.io/ReportGenerator) CLI tool:

```bash
dotnet tool install --global dotnet-reportgenerator-globaltool
```

- Main command

```bash
reportgenerator "-reports:test/*/*.opencover.xml" "-targetdir:coverage" "-reporttypes:Html"
```

Run the command in the **root** project.

- Coverage report

Find the folder **coverage** and open **index.html**.

## Improvements

- Search endpoint cache
- Performance query for search (remove skip and limit)
