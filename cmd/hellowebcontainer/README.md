# Generates a web app container suitable for deployment on Google Cloud Run

## Setup
Note: Google Cloud Shell already include skaffold, docker, kubectl etc.

For local installation:
```
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
install skaffold /home/siuyin/bin/
```

## Build
1. skaffold build
1. note the git commit (first 7 characters of git log)
1. specify the tag with the Cloud Run image. See https://console.cloud.google.com/run