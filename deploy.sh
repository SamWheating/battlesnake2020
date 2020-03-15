#!/bin/bash

SERVICE_NAME=battlesnake
IMAGE=gcr.io/personalsite-264919/battlesnake:latest
PROJECT=personalsite-264919

gcloud config set account samwheating@gmail.com
docker build -t $IMAGE .
docker push $IMAGE
gcloud run deploy $SERVICE_NAME --image $IMAGE --region us-central1 --platform managed --project $PROJECT