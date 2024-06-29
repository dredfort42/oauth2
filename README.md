# OAuth 2.0 Authorization Service

## Introduction

A simple and fast OAuth 2.0 Authorization Service written in Go. This service provides authorization to access protected resources, based on the OAuth 2.0 Authorization Framework ([RFC6749](https://tools.ietf.org/html/rfc6749)) and the OAuth 2.0 Device Authorization Grant ([RFC8628](https://tools.ietf.org/html/rfc8628)).

The service use Postgres as a database for storing user data, sessions, and devices. It is designed for use in a microservice architecture and is ready for horizontal scaling.

## Features

- OAuth 2.0 Authorization
- OAuth 2.0 Device Authorization

## Configuration

The service is configured using a `config.ini` file located in the root folder of the project. The configuration file has the following structure:

```ini
# Auth service configuration file

[auth] # Auth service configuration
# The host and port that the auth service will bind to and listen on for incoming connections from clients and other services
auth.host=localhost
auth.port=4242
# Device verification URL that will be sent to the user's device for verification
auth.device.verification.url=https://localhost:4242/auth/device/verify
# CharSet for device verification code generation. Can be any letters, numbers or symbols you want
auth.device.verification.code.charset=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
# Length of the device verification code
auth.device.verification.code.length=8
# Expiration time for the device verification code in seconds
auth.device.verification.code.expiration=300
# Frequency of polling for the device tokens in seconds
auth.device.verification.code.attempts=3

[jwt] # JWT configuration
# Secret key
jwt.secret=rupl_auth_secret
# Expiration time for the JWT tokens in seconds
jwt.onetime.access.token.expiration=1800    # 30 minutes
jwt.onetime.refresh.token.expiration=86400  # 1 day
jwt.browser.access.token.expiration=3600    # 1 hour
jwt.browser.refresh.token.expiration=604800 # 7 days
jwt.device.access.token.expiration=86400    # 1 day
jwt.device.refresh.token.expiration=2592000 # 30 days

[database] # Database configuration
# Postgres database connection parameters
db.host=localhost
db.port=5432
db.security.ssl=disable
db.user=user_name
db.password=user_password
# Database cleanup interval in seconds when the service will remove expired sessions and devices (0 - disable cleanup)
db.cleanup.interval=3600 # 1 hour
# Database name and tables
db.database.name=auth_service
db.table.users=auth_users
db.table.sessions=auth_sessions
db.table.devices=auth_devices
```

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
