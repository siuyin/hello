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
