# Lendo Backend

## API
The backend API was developed with [the Gin web framework](https://github.com/gin-gonic/gin), and consists of three endpoints:

* POST /word

Adds a new word to the data structure

* POST /synonym/:word

Adds a new synonym (provided in the request body) to the word defined in the URL path

* GET /synonyms/:word

Returns all synonyms for the word defined in the URL path

## Data structure
A graph data structure was used to store the words and their synonyms. This was chosen to efficiently handle the cyclical relationship between the words and their synonyms.

## Unit tests
Unit tests have been written to cover the API's major functionality. All tests can be run via `go test ./...` from the project's root directory.

## Running locally
The API can be run locally via `go run main.go` from the project's root directory. This will serve the API at `localhost:8080`.

## Public Deployment
The API has been deployed publicly to GCP's Cloud Run, and is available at [https://lendo-backend-api-k3tawfx6ta-ew.a.run.app](https://lendo-backend-api-k3tawfx6ta-ew.a.run.app). Deployment is handled by the `cloudbuild.yaml` file in the project's root directory.