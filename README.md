# GoBackendProject

## Project Summary:
* This is a Go backend project for training
* It is a Greetings API, with CRUD methods
* Available Endpoints:
    * GET /healthcheck
        * To smoketest the API
    * GET /greetings
        * Returns all Greetings stored in the API
    * GET /greeting/`:uuid`
        * Returns a specific Greeting, where `:uuid` is the Greeting ID
    * GET /greeting/random
        * Returns a random Greeting
    * POST /greeting
        * Adds a new Greeting to the stored Greetings
    * PUT /greeting/`:uuid`
        * Modifies the Greeting message where `:uuid` is the Greeting ID
    * DELETE /greeting/`:uuid`
        * Deletes the Greeting where `:uuid` is the Greeting ID
* Bespoke Middleware:
    * Logging
        * Provides useful logs when requests are made and finished to allow DevOps to diagnose bugs and set up alerts
    * Authentication
        * Ensures only authenticated users can modify the system / endpoints that aren't read-only are protected

## To clone and run this API locally you must:
1. First you will require: `go v1.24.0` with this version as minimum requirement
2. Clone this repo using command `git clone https://github.com/EliR94/GoBackendProject.git`
4. Run `go run .` to run the API locally, and use a tool such as Postman use the locally hosted API (using header `X-Auth` : `mySecretCodeABC123` for POST/PUT/DELETE requests)

## API Documentation:
* GET /healthcheck
    * Request:
        * No request body or headers required
    * Expected Response:
        * Status Code: `200`
        * Body: `"All Good!"` - if the API is running without error
* GET /greetings
    * Request:
        * No request body or headers required
    * Expected Response:
        * Status Code: `200`
        * Body: `{
  "items": [
    {
      "id": "12345678-9012-3456-7890-123456789012",
      "message": "Hello World"
    },
    ...
  ]
}`
* GET /greeting/`:uuid`
    * Request:
        * No request body or headers required
        * `:uuid` is the Greeting ID
    * Expected Response:
        * Status Code: `200`
        * Body: `{
    "id": "12345678-9012-3456-7890-123456789012",
    "message": "Hello World"
}`
* GET /greeting/random
    * Request:
        * No request body or headers required
    * Expected Response:
        * Status Code: `200`
        * Body: `{
    "id": "12345678-9012-3456-7890-123456789012",
    "message": "Random Greeting Here"
}`
* POST /greeting
    * Request:
        * Requires request body `{ "Message" : "Hello World" }` and a authorised X-Auth header
    * Expected Response:
        * Status Code: `201`
        * Body: `{
  "id": "12345678-9012-3456-7890-123456789012",
  "message": "Hello World"
}`
* PUT /greeting/`:uuid`
    * Request:
        * Requires request body `{ "Message" : "New Message Here" }` and a authorised X-Auth header
        * `:uuid` is the Greeting ID
    * Expected Response:
        * Status Code: `200`
        * Body: `{
    "id": "12345678-9012-3456-7890-123456789012",
    "message": "New Message Here"
}`
* DELETE /greeting/`:uuid`
    * Request:
        * Requires a authorised X-Auth header
        * `:uuid` is the Greeting ID
    * Expected Response:
        * Status Code: `204`
        * No Response Body
