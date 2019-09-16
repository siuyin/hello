FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p bin && CGO_ENABLED=0 go build -o bin -v ./cmd/helloweb && ls -lF bin

FROM scratch
COPY --from=builder /app/bin/helloweb /app
EXPOSE 8080
ENV SUBJECT="World in Docker"
CMD ["/app"]
