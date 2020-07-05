#!/bin/bash
set -e

gcloud config set account samwheating@gmail.com

gcloud app deploy --quiet