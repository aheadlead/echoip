# Simple Go HTTP Server

This is a simple HTTP server written in Go. It will return the IP address of any clients who send a GET request to the server. Also, it supports Bearer token verification which can be defined via an environment variable.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.17 or later
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Building and Running

To build and run this application, use Docker Compose:

```bash
docker-compose up --build
```

## Environment Variables

The application relies on two required environment variables:

* `BEARER_TOKEN` - This should be your authorization bearer token.
* `PORT` - The port number on which you want the server to listen.

You can define these variables in an .env file.

## Versioning

We use [SemVer](https://semver.org/) for versioning.

## Authors

ChatGPT

## License

This project is licensed under the MIT License - see the [LICENSE.md](http://license.md/) file for details.


