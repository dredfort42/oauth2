# OAuth 2.0 Authorization Service

## Introduction

A simple and fast OAuth 2.0 Authorization Service written in Go. This service provides authorization to access protected resources, based on the OAuth 2.0 Authorization Framework ([RFC6749](https://tools.ietf.org/html/rfc6749)) and the OAuth 2.0 Device Authorization Grant ([RFC8628](https://tools.ietf.org/html/rfc8628)).

The service use Postgres as a database for storing user data, sessions, and devices. It is designed for use in a microservice architecture and is ready for horizontal scaling.

## Features

- OAuth 2.0 Authorization
- OAuth 2.0 Device Authorization

## Configuration

The service is configured with a [config.ini](config.ini) file found at `/app/config.ini` or another file specified using the --config flag.

## Installation

To install the service, you need to have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/](https://golang.org/).

After installing Go, clone the repository and build the service using the following commands:

```bash
git clone https://github.com/dredfort42/oauth2.git
cd oauth2
go build -o auth ./cmd/auth/main.go
```

## Usage

### Running the service

To start the service, run the following command:

```bash
./auth
```

The service will start and listen on the host and port specified in the configuration file.

### Running the service in Debug mode

To run the service in Debug mode, set the DEBUG environment variable to true before starting the service:

```bash
env DEBUG=true ./auth
```

The service will start in Debug mode and print additional information to the console.

### Running the service with a specific configuration file

To run the service with a specific configuration file, set the --config flag with the path to the configuration file while starting the service:

```bash
./auth --config /path/to/my_config.ini
```

### Running the service in Docker

To run the service in Docker, build the Docker image using the following command:

```bash
docker build -t auth .
```

After building the Docker image, run the service using the following command:

```bash
docker run -p 4242:4242 auth
```

## API

The service provides API endpoints for OAuth 2.0 Authorization and OAuth 2.0 Device Authorization. The API endpoints are described in the [openapi.yaml](/api/openapi.yml) file.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
