# Generates a web app container suitable for deployment on Google Cloud Run

1. skaffold build
1. note the git commit (first 7 characters of git log)
1. specify the tag with the Cloud Run image. See https://console.cloud.google.com/run