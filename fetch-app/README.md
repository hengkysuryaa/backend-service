# Fetch App

## Prerequisite

- Golang 1.16.x or newer
- Docker
- Docker compose

## Quick start

**Important Note**

Please create `.env` file in this directory. Format:
```
REST_PORT=:8080
JWT_SIGNATURE_KEY=<jwt secret>
RESOURCE_URL=https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list
CURRENCY_EXCHANGE_URL=https://api.apilayer.com/exchangerates_data
CURRENCY_EXCHANGE_API_KEY=<api key from https://exchangeratesapi.io/>
```

**For development using docker-compose**, please run below command
```
docker-compose up --build
```

**For local development, without docker**, please run

To install all golang dependencies in `go.mod` file
```
go mod download
```
To start the server
```
go run main.go
```
The service will be started on port :8080 (or change it via `.env` file and `docker-compose.yml`)

## API Contract
Please visit this API documentation for list of avaiable APIs and the contract
https://documenter.getpostman.com/view/23632461/2s83tAsZin#f7d66214-3522-4361-a47e-3607bc75a038

## Project folder structure
- cmd: store commands of the app to run, ex: run rest server
- internal
    - domain: define entity, dto (models), repository layer, usecase (service) layer
    - ports: handle incoming request, ex: rest (http)
- pkg: define custom package that can be used anywhere in project