FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p bin && CGO_ENABLED=0 go build -o bin ./cmd/secureweb 

FROM scratch
COPY --from=builder /app/bin/secureweb /app
EXPOSE 8080
CMD ["/app"]

# vim: set filetype=dockerfile :
