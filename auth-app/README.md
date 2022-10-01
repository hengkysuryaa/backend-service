# Auth App

## Prerequisite

- NodeJS v16.x or above
- Docker
- Docker compose
- SQLite3

## Quick start

**Important Note**

Please create `.env` file in this directory. Format:
```
API_PORT=8081
JWT_SECRET=<jwt secret>
```

**For development using docker-compose**, please run below command
```
docker-compose up --build
```

**For local development, without docker**, please run

To install app packages
```
npm install
```
To start the server
```
npm start index.js
```
The service will be started on port :8081 (or change it via `.env` file and `docker-compose.yml`)

## API Contract
Please visit this API documentation for list of avaiable APIs and the contract
https://documenter.getpostman.com/view/23632461/2s83tAsZin#93722c1f-d6c5-44cd-ba50-f49774a08d78

## Project folder structure
- controllers: handle request-response from endpoint
- db: store file-based database and the ORM connection
- library: define custom functions that can be used anywhere in project
- models: define object in database table
- routes: define app route / endpoint and middleware
- services: apply business logic and validation