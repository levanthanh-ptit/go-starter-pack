# Go Starter Pack

An boilerplate for golang project that boot you up.

## Features

- Using DDD Pattern as:
  - Application layer: `transport`
  - Domain layer: `domain`, `services`
  - Infrastructure layer: `provider`, `repository`

## Project Structure

```sh
.
├── cmd
│   └── app
│       └── main.go
├── configs
│   ├── storage.go
│   ├── server.go
│   ├── config.go
│   └── auth.go
├── internal
│   ├── transport       # Contain server interface and HTTP transport implement
│   │   ├── server.go
│   │   └── http
│   │       └── ...
│   ├── services        # Contain domain services
│   │   ├── user.go
│   │   └── auth.go
│   ├── domain          # Contain business models and service interfaces
│   │   ├── user.go
│   │   └── auth.go
│   ├── provider        # Contain interfaces and implements for library provider
│   │   ├── token_provider.go
│   │   ├── password_hash_provider.go
│   │   └── jwt
│   └── repository      # Contain repository interfaces and implements
│       ├── user.go
│       └── gormrepo
│           └── ...
├── tests               # Contain integration tests
│   └── ...
├── go.sum
├── go.mod
├── docker.env
├── docker-compose.yml
├── REQUIREMENT.md
├── README.md
├── Makefile
├── LICENSE
└── Dockerfile
```

# Instructions

Make sure you have Go installed ([download](https://go.dev/dl/)). Version `1.17` or higher is required.

Make sure you have Docker installed ([instructions](https://docs.docker.com/engine/install/)).

Make sure you have `make` installed for running the scripts.

For Linux:

```sh
sudo apt install make
```

For MacOS:

```sh
brew install make
```

<br/>

## Start Server

Using command bellow to build and run on Docker Compose

```sh
make start
```
- `Togo` app will available on `127.0.0.1:4000`
- `Redis` will available on `127.0.0.1:56379`
- `PostgreSQL` will available on `127.0.0.1:55432`

<br/>

## Run Unit Tests

Using command bellow to run unit tests

```sh
make run-unit-test
```

<br/>

## Run Integration Tests

Make sure Togo server build and start successfully then use the command bellow to run integration tests.

```sh
make run-unit-test
```

<br/>

# License

MIT