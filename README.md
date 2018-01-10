# Advert service for vertical classified

## Features:
* Completely ready for use
* Specification first
* Code generation for boilerplate code
* Auto-migration for DB
* CQRS + ES
* Swagger
* Docs [redoc](https://github.com/Rebilly/ReDoc)

## Dependency
* [Docker](https://www.docker.com/)
* [Golang >= 1.9](https://golang.org/)
* [dep](https://github.com/golang/dep)
* [Go-swagger](https://github.com/go-swagger/go-swagger)

## Setup
After install all dependencies and update `swagger.yml` just run command:
```
$ make run
```

## Documentations
```
$ make docs
```

## Testing
```
$ make test
```

## Architecture
![architecture](https://raw.githubusercontent.com/kot13/vertigo/true-purpose/static/architecture.png)