# go-api-gateway

This is a simple API gateway written in GO which can be used to redirect requests to other services. This gateway utilises gRPC to communicate with the services hence the `.proto` files.

The services which can be called are as follows:

- [Auth Service](https://github.com/Bryan-BC/go-auth-microservice)
- [Course Service](https://github.com/Bryan-BC/go-course-microservice)
- [Timetable Service](https://github.com/Bryan-BC/go-timetable-microservice)

## Prerequisites

Since this gateway uses gRPC, it requires the [protocol buffer compiler](https://grpc.io/docs/protoc-installation/) installed. Additionally, this repository uses a Makefile to compile the `.proto` files and run the gateway itself. Hence, it requires [make](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows) to be installed.

## Setup

To run the gateway, simply cd into the repository and run the following command:

`make start`
