# Hello world examples and code templates

## Testing

```
go test ./...
```

## Building binaries

The folowing only works with go 1.13 onward.
```sh
mkdir bin
go build -o bin ./...
```

## Package as docker container

```sh
docker build --tag siuyin/junk --file helloworld.dockerfile .
docker build -t siuyin/junk -f helloworld.dockerfile .
docker build -t siuyin/junk -f goodbyeworld.dockerfile .
```

## Deloy to kubernetes with skaffold

```sh
export KUBECONFIG=/path/to/kube/config
skaffold run               # deploys dev release
skaffold run -p prod       # deploys prod release
```
This currently deploys helloweb with opencensus monitoring.

## Example of two pod deployment to kuberenetes
```sh
export KUBECONFIG=/path/to/kube/config
skaffold run -f hello-nats-stream.skaffold.yaml
skaffold run -f hello-nats-stream.skaffold.yaml -p prod
```

## Example with kustomize and ticktock binary
```sh
export KUBECONFIG=/path/to/kube/config
skaffold run -f ticktock.skaffold.yaml # deploys base kustomization
skaffold run -f ticktock.skaffold.yaml -p prod # deploys overlay/prod kustomization
skaffold run -f ticktock.skaffold.yaml -p test # deploys overlay/test kustomization

```
## keycloak.kustomize
keycloak.kustomize holds kustomization scripts for
my attempt at a production ready keycloak installation.

To install / delete:
```sh
export KUBECONFIG=/path/to/kube/config
kustomize build keycloak.kustomize/overlays/prod | kubectl apply -f -

kustomize build keycloak.kustomize/overlays/prod | kubectl delete -f -
```

## keycloak sercure web app
Deploy keycloak server as shown above.
Get the service address:
```sh
export KUBECONFIG=/path/to/kube/config
kubectl get svc -l sys=keycloak

skaffold run -f secureweb.skaffold.yaml -p dev
```
```

