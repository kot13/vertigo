# Advert service for vertical classified

## Features:
* Completely ready for use
* Specification first
* Code generation for boilerplate code
* Auto-migration for DB
* CQRS

## Dependency
* [Docker](https://www.docker.com/)
* [Golang >= 1.9](https://golang.org/)
* [gvt](https://github.com/FiloSottile/gvt)
* [PostgreSQL](https://www.postgresql.org/)

## Setup
After install all dependencies, follow these simple steps: 
1. Update specs in `./spec.yml`
2. Update config for connect to DB 
3. Create DB scheme:
```
$ make migrate
```
4. Start app:
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