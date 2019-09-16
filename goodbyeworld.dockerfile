FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o bin -v ./cmd/goodbyeworld && go version

FROM scratch
COPY --from=builder /app/bin/goodbyeworld /app
CMD ["/app"]
