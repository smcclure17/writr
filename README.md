# Writr

Backend for Writr, a Google Docs/Notion/Dropbox Paper clone.

## Setup

### AWS DynamoDB/Database

Writr uses AWS DynamoDB for persisted database storage.

#### AWS Config

Set up AWS configuration to the project you want to use with `aws configure`.
To create a configuration, create an IAM role with the permissions `AmazonDynamoDBFullAccess`.
See the docs for help with [creating a new role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html).

#### Creating Necessary Tables

To create the necessary tables, we use Terraform IaaS.
Construct the tables by running `terraform apply` (with your correct AWS config).

### Redis

Writr uses Redis as an in-memory cache for fast/immediate document recovery.

Install Redis (on mac) with `brew install redis`.

## Running Locally

We can run the Writr service (mostly) locally for development purposes.

(I have not set up a local DynamoDB workflow, so all DB connections will connect to the production database).

1. Start Redis by running `redis-server`
2. Start the service by running `go run app/websockets.go`
3. Explore/test the service by navigating to `http://localhost:8081/`

## TODOs

Things to fix/update:

- Add env variables instead of hard coding values (e.g. port numbers)
- Improve error handling (see ["Go's Error Handling Is a Form of Storytelling"](https://preslav.me/2023/04/14/golang-error-handling-is-a-form-of-storytelling/))
- Use Enums for categorical values (e.g. table names)
- Containerize and deploy (Docker + Docker compose)
- Scaling (use a message queue behind load-balanced servers)
- Convert these to Github tickets/issues
