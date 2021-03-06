# band-app-go

No purpose API app to strengthen my skills with Go, Docker, MongoDB, and API building

## Getting started

These directions will guide you to getting the project setup, modifying, and running the band-app.

### Prerequisites

* GOLANG 1.16+
* git
* docker

### Nice to haves

* Postman
* MongoDB Compass

## Building the binary

From cmd/band-app-server run the following command.

```bash
go build
```

## Makefile

Look at the Makefile for a list of commands in regards to running/building the project.

## Profiling the project

With the service running you can visit the /debug/pprof/ endpoint to view a multitude of profiling features.

For any generated profiles you can use this command to start a web server to view them.

```bash
go tool pprof -http=127.0.0.1:6060 <profile>
```

For any generated traces you can use this command to start a web server to view them.

```bash
go tool trace <trace>
```

## Brew osx package manager

Can be found here with instructions on how to install <https://brew.sh/>

## golang

Can be found here <https://golang.org/>

Installing with brew:

```bash
brew install go
```
