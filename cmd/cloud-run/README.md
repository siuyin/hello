# cloud-run try

Trying out Google's cloud run.

## Deployment

1. Create a docker image and push it to a registry.
   skaffold build creates image siuyin/junk:{RELEASE} where RELEASE is from the environment.

```
. release.env
skaffold build
```

1. Deploy the image as a service on google cloud-run as follows:

```
gcloud run deploy foo --image=siuyin/junk:v1 --cpu=100m --memory=64Mi --platform=managed --region=us-west1 --port=8080 --timeout=5s
```
