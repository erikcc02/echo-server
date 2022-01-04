FROM golang:1.17.5-stretch AS build
WORKDIR /app
COPY . /app
RUN go test ./...
RUN go build -o echoserver ./src/cmd/main.go

FROM golang:1.17.5-stretch AS runtime
WORKDIR /echoapp
COPY --from=build /app/echoserver .
RUN chmod 777 ./echoserver
EXPOSE $ECHOSERVER_PORT
CMD [ "./echoserver" ]