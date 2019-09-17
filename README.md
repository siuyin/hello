# Hello world examples and code templates

## Testing

```
go test ./...
```

## Building binaries

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

## Example of two pod deployment to kuberenets
```sh
export KUBECONFIG=/path/to/kube/config
skaffold run -f hello-nats-stream.skaffold.yaml
skaffold run -f hello-nats-stream.skaffold.yaml -p prod
```
