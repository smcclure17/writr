# Writr

Backend for a Google Docs/Notion/Dropbox Paper clone.

## Setup

### Redis

Writr uses Redis as an in-memory cache for fast/immediate document recovery.

Install Redis (on mac) with `brew install redis`.

## Running Locally

We can run the Writr service (mostly) locally for development purposes.

(I have not set up a local DynamoDB workflow, so all DB connections will connect to the production database).

1. Start Redis by running `redis-server`
2. Start the service by running `go run app/websockets.go`

## Deployment

Set necessary environment variables and run `./deploy.sh` to build a docker image and deploy
the image to Google Cloud run.

## TODOs

Things to fix/update:

- Add env variables instead of hard coding values (e.g. port numbers)
- Improve error handling (see ["Go's Error Handling Is a Form of Storytelling"](https://preslav.me/2023/04/14/golang-error-handling-is-a-form-of-storytelling/))
- Use Enums for categorical values (e.g. table names)
- Containerize and deploy (Docker + Docker compose)
- Scaling (use a message queue behind load-balanced servers)
- Convert these to Github tickets/issues
