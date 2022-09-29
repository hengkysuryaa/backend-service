# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

COPY . /go/src/github.com/hengkysuryaa/backend-service/fetch-app
WORKDIR /go/src/github.com/hengkysuryaa/backend-service/fetch-app

RUN go mod download

RUN go build -o /fetch-app

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /fetch-app /fetch-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/fetch-app"]