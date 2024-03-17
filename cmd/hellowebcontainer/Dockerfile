FROM golang:1.22.1 as build
WORKDIR /src
COPY . /src/
RUN go build -o /bin/app ./main.go


FROM gcr.io/distroless/base-nossl-debian12:nonroot
EXPOSE 8080
COPY --from=build /bin/app /bin/app
COPY --from=build /src /src
CMD ["/bin/app"]
