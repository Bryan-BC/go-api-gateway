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

Once this is run, the gateway will be running on port 3000.

## Endpoints

- POST `/auth/register` - Register a new user. The request body should be a JSON object similar to the following:

```json
{
  "username": "Bece",
  "password": "thequickbrownfoxjumpedoverthelazydog"
}
```

- POST `/auth/login` - Login a user. The request body should be a JSON object similar to the following:

```json
{
  "username": "Bece",
  "password": "thequickbrownfoxjumpedoverthelazydog"
}
```

- GET `/course/{id}` - Get a course by ID.

- POST `/course` - Create a new course. The request body should be a JSON object similar to the following:

```json
{
  "course": {
    "id": "MH1301",
    "name": "Discrete Mathematics",
    "description": "Teaches the basics of discrete mathematics",
    "schedule": [
      {
        "name": "Discrete Mathematics",
        "index": "03669",
        "days": [
          {
            "day": "Monday",
            "timings": [
              {
                "start": 10,
                "end": 13
              }
            ]
          },
          {
            "day": "Wednesday",
            "timings": [
              {
                "start": 18,
                "end": 19
              }
            ]
          }
        ]
      }
    ]
  }
}
```

- DELETE `/course/{id}` - Delete a course by ID.

- GET `/timetable/{id}` - Get a timetable by ID.

- POST `/timetable` - Create a new timetable. The request body should be a JSON object similar to the following:

```json
{
  "courses": ["MH1300", "MH1301"]
}
```

All the above endpoints are relative to the base URL of the gateway (Default: `http://localhost:3000`).

The endpoints prefixed with `/course` and `/timetable` require a bearer token which is generated as the response of the `/auth/login` endpoint.
