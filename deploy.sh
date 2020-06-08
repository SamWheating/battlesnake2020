#!/bin/bash
set -e


SERVICE_NAME=battlesnake
IMAGE=gcr.io/personalsite-264919/battlesnake:latest
PROJECT=personalsite-264919

gcloud config set account samwheating@gmail.com
docker build --no-cache -t $IMAGE .
docker push $IMAGE

# Deploy to cloud run in us-west-1 (The same zone as the battlesnake gameserver)
gcloud run deploy $SERVICE_NAME --image $IMAGE --region us-west1 --platform managed --project $PROJECT