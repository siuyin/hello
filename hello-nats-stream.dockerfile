FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p bin && CGO_ENABLED=0 go build -o bin -v ./cmd/hello-nats-stream && ls -lhF bin

FROM scratch
COPY --from=builder /app/bin/hello-nats-stream /app
CMD ["/app"]

# vim: set filetype=dockerfile :
