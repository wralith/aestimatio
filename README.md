# Aestimatio

Aestimatio is a productivity app!

## Setup

This section will be updated after i wrote dockerfiles and implement docker-compose.  
For now:

```sh
# Start Frontend
cd client/web
pnpm dev
# Start Services
cd server/{service}
go run cmd/main.go
```

REST docs would be at `http:localhost:8080/docs`

## Techs

### Server Side

- Three services written in `go`.
- I tried to apply `DDD` and hexagonal architecture in services.
- `protobuf` and `gRPC` based communication between services, REST api-gateway to communicate with outside.
- Separate auth service deals with authentication and passes credentials to other services via api-gateway.
- Api-gateway implemented with `Echo`.
- Nearly all of the business logic have unit tests.
- I followed `repository` pattern, any database will be introduced as an adaptor easily.
- Docs for REST api created with [swaggo/swag](https://github.com/swaggo/swag) and [rapidoc](https://github.com/rapi-doc/RapiDoc).

### Client Side

- React...
- React Query because of utils it brings such as caching, and getting state of the mutation, query etc.
- [Mantine](https://github.com/mantinedev/mantine) as a component library because it looks cool...
- [Zustand](https://github.com/pmndrs/zustand) for global state management (i.e. auth).

## Todos

- Error Messages are awful, need to be fixed
- Calendar
- Daily Tasks

## Far far away Todos

- Social media like groups, tasks with multiple users, share achievement, share statistics, chat etc.
