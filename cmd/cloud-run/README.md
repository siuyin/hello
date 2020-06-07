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

## Google specifics

### Authorized docker within cloud console

gcloud auth configure-docker

### Creating a service account to access the registry

gcloud iam service-accounts create junk-account

gcloud projects add-iam-policy-binding PROJECT_ID --member "serviceAccount:NAME@PROJECT_ID.iam.gserviceaccount.com" --role "roles/ROLE"

eg.
gcloud projects add-iam-policy-binding qwiklabs-gcp-04-487598f65dc1 --member "serviceAccount:junk-account@qwiklabs-gcp-04-487598f65dc1.iam.gserviceaccount.com" --role "roles/storage.admin"

For roles, see: https://cloud.google.com/container-registry/docs/access-control

### Creating a key file for access credentials

gcloud iam service-accounts keys create keyfile.json --iam-account [NAME]@[PROJECT_ID].iam.gserviceaccount.com

eg.
gcloud iam service-accounts keys create keyfile.json --iam-account junk-account@qwiklabs-gcp-04-487598f65dc1.iam.gserviceaccount.com

### Using the key file

cat keyfile.json | docker login -u _json_key --password-stdin https://[HOSTNAME]

cat secret.keyfile.json | docker login -u _json_key --password-stdin https://gcr.io 

(gcr.io, us.gcr.io or asia.gcr.io, ...)


### Building and pushing with skaffold

skaffold -d gcr.io/qwiklabs-gcp-04-487598f65dc1 build

And in cloud console:

gcloud run deploy foo --image=gcr.io/qwiklabs-gcp-04-487598f65dc1/siuyin_junk:v1 --cpu=1 --memory=64Mi --platform=managed --region=us-west1 --port=8080 --timeout=5s --allow-unauthenticated
