FROM golang:1.22.1 as build
WORKDIR /src
COPY . /src/
RUN ls -l
RUN go build -o /bin/app ./main.go


FROM scratch
EXPOSE 8080
COPY --from=build /bin/app /bin/app
CMD ["/bin/app"]