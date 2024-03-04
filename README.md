# TTMI-SG Golang Backend Challenge

Indego is Philadelphia's bike-sharing program, with many bike stations in the city.

The Indego GeoJSON station status API provides a realtime snapshot of the number of bikes available, number of open docks available (not currently containing a bike), and total number of docks at every station. This API is free and requires no API key.

The Open Weather Map API provides a realtime snapshot of the current weather in a given city. Since Philadelphia is a small geographical area it is sufficient to obtain the weather for a geographical location central to Philadelphia. This API has a free plan, you will need to sign up for an API key.

Using Golang, create a new API server which accumulates data over time and provides access to historical data for both weather and Indego bike availability.

## API-endpoints

Use a static token and protect all the endpoints. If that static token is not provided or is invalid return error response with relevant HTTP status code.
