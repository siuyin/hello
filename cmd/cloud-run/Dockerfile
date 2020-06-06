FROM golang:1.14-buster as build
WORKDIR /src
COPY go.mod go.sum /src/
RUN go mod download
ADD . /src
RUN go build -o /build/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /build/app /
CMD ["/app"]
