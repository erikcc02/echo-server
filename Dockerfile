FROM golang:1.17.5-stretch AS build
WORKDIR /app
COPY . /app
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o echoserver ./src/cmd/main.go

FROM alpine:3.15 AS runtime
WORKDIR /echoapp
COPY --from=build /app/echoserver .
RUN chmod 777 ./echoserver
EXPOSE 8000
CMD [ "./echoserver" ]