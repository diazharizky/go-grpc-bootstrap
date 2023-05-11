# gRPC Server Bootstrap Code using Go

This template could reduce effort of creating new gRPC server using Go programming language.

## Project Structure

The structure is actually inspired by <https://github.com/golang-standards/project-layout>, nothing much changed on this project.

### `services` Package

A service is responsible to handle a single task/business process or could be a set of other services.

### App Context

Is a core component that includes services, repositories, databases, cache, etc..

## Getting Started

Before everything else, if you are unable to find `pb` package kindly do the following

```bash
make protogen
```

To start the server

```bash
make run
```

To run unit tests

```bash
make test
```
