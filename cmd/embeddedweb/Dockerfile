FROM golang:1.22.1 as build
WORKDIR /src
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/app ./main.go


FROM gcr.io/distroless/static-debian12:nonroot
EXPOSE 8080
COPY --from=build /bin/app /bin/app
COPY --from=build /src /src
CMD ["/bin/app"]
