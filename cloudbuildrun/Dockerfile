FROM golang:1.23 AS build
WORKDIR /src
COPY ./* .
RUN CGO_ENABLED=0 go build -o /bin/app ./*.go

FROM scratch
COPY --from=build /bin/app /bin/app
CMD ["/bin/app"]