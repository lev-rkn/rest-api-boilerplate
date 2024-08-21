# rest-api-service

## Navigation
1. [Description](#description)
1. [API](#api)
1. [Tests](#tests)

## Description
Run:
```bash 
docker-compose up -d
```

## Migrations
Run migrations up before use this service:
```bash
make migrations-up
```

Create new migration file:
```bash
goose -dir migrations create {{migration_name}} sql
```

## API
http server address by default: localhost:8080  
Swagger docs: <http://localhost:8080/swagger/>

## Tests
```bash
make cover
```
