#!/bin/bash


# usage: ./deploy.sh [--vpc] 
if [ "$1" == "--help" ]; then
    echo "usage: ./deploy.sh [--vpc]"
    echo "  --vpc: set to create the VPC connector"
    exit 0
    fi

if [ -z "$REDIS_URL" ]; then
    echo "${REDIS_URL}"
    echo "REDIS_URL not set. Please set the REDIS_URL environment variable to the IP of the Redis instance."
    exit 1
    fi


# if command line arg --vpc is passed, then create a VPC connector
if [ "$1" == "--vpc" ]; then

    if [ -z "$VPC_IP_RANGE" ]; then
    echo "VPC_IP_RANGE not set."

    echo "Creating VPC connector..."
    gcloud compute networks vpc-access connectors create redis-connector \
    --network default \
    --region us-central1 \
    --range $VPC_IP_RANGE
    fi


echo "Building and submitting container image..."
gcloud builds submit --tag gcr.io/writr-dev-384017/writr-dev

echo "Deploying container image..."
gcloud run deploy --image gcr.io/writr-dev-384017/writr-dev \
--platform managed \
--allow-unauthenticated \
--region us-central1 \
--vpc-connector projects/writr-dev-384017/locations/us-central1/connectors/redis-connector \
--set-env-vars REDIS_URL=$REDIS_URL \
--port 8081

echo "Done!"