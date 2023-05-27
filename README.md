# gRPC App Bootstrap Code using Go

This template could reduce effort of creating new gRPC app using Go programming language.

## Project Structure

The structure is mostly inspired by <https://github.com/golang-standards/project-layout>, nothing much changed on this project.

### `services` Package

A single service is responsible to handle a single task/business process or could be a set of call of other services.

### App Context

Is a core and shared component that includes services, repositories, databases, cache, etc..

## Getting Started

Before everything else, if you are unable to find `pb` package kindly do the following command

```bash
make protogen
```

To start the server

```bash
make run
```

To run the unit tests

```bash
make test
```
