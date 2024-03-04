# TTMI-SG Golang Backend Challenge

Indego is Philadelphia's bike-sharing program, with many bike stations in the city.

The Indego GeoJSON station status API provides a realtime snapshot of the number of bikes available, number of open docks available (not currently containing a bike), and total number of docks at every station. This API is free and requires no API key.

The Open Weather Map API provides a realtime snapshot of the current weather in a given city. Since Philadelphia is a small geographical area it is sufficient to obtain the weather for a geographical location central to Philadelphia. This API has a free plan, you will need to sign up for an API key.

Using Golang, create a new API server which accumulates data over time and provides access to historical data for both weather and Indego bike availability.

## API-endpoints

Use a static token and protect all the endpoints. If that static token is not provided or is invalid return error response with relevant HTTP status code.

### Store snapshot data

An endpoints which downloads fresh data from Indego GeoJSON station status API and Open Weather Map API then stores it inside PostgreSQL.

```
# this endpoint will be trigger every hour to fetch the data and insert it in the PostgreSQL database
POST http://localhost:3000/api/v1/data-fetch-and-store-it-db
```

### Snapshot of all stations at a specified time

Data for all stations as of 10am Universal Coordinated Time on September 1st, 2019:

```
GET http://localhost:3000/api/v1/stations?at=2019-09-01T10:00:00Z
```

This endpoint should respond as follows, with the actual time of the first snapshot of data on or after the requested time and the data:

```
{
at: '2019-09-01T10:00:00Z',
stations: { /_ As per the Indego API _/ },
weather: { /_ As per the Open Weather Map API response for Philadelphia _/ }
}
```

### Snapshot of one station at a specific time

Data for a specific station (by its kioskId) at a specific time:

```
GET http://localhost:3000/api/v1/stations/{kioskId}?at=2019-09-01T10:00:00Z
```

The response should be the first available on or after the given time, and should look like:

```
{
  at: '2019-09-01T10:00:00Z',
  station: { /* Data just for this one station as per the Indego API */ },
  weather: { /* As per the Open Weather Map API response for Philadelphia */ }
}
```

Include an at property in the same format indicating the actual time of the snapshot.

If no suitable data is available a 404 status code should be given.

## Unit tests

Write functional and API test the tool of your choice.

## Hosting details

You will need to make your API available on a server running in local environment.

## Criteria

Your work will be evaluated primarily on:

- README file how to run/test in local
- API documentation
- Unit tests for API endpoints and also do functional testing
- Proper error handling
- Efficient Postgres queries and indexes
- How to submit your work
- TTMI will prepare your repository for submit, TTMI-assessments-{your_github_account}

## Extra credit

- oAuth2 integration to protect your APIs by registering Auth0 free account.
- Host this backend application on AWS
- A simple front end React application offering a visualization of all or part of the data utilizing the API you have built as a back end.
- Anything else you think is cool, relevant, and consistent with the other requirements.

## Pre requisities

- Install the docker environment
- Setup the `app.env`

## How to run in local

First of all, build the docker image for postgresSQL

```
docker-compose up -d
```

Second, create a table by migrating data structure into postgreSQL

```
go run migrate/migrate.go
```

Third, run in local

```
go run main.go
```
