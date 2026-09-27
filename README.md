# GoDispatch

GoDispatch is a generic dispatch and matching platform built with Go.

The project is designed to support multiple domains such as Taxi and Food Delivery using a reusable Universal Dispatch Model (UDM).

## Current Architecture

```text
Taxi API ──┐
           ├── UDM ── Spatial
Food API ──┘
              │
           Adapters
```

## Technology

* Go
* REST API
* PostgreSQL
* Redis
* NATS
* Docker

## Project Structure

```text
go-dispatch/
├── cmd/
│   ├── taxi-api/
│   └── food-api/
│
└── internal/
    ├── kernel/
    ├── udm/
    ├── spatial/
    ├── taxi/
    ├── food/
    └── adapters/
```

## Status

🚧 Under Development
