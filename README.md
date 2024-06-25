# kelvin
Kelvin challenge

## Run Locally

```bash
  docker-compose up --build
```

## Features
- HTTP Rest endpoint on :8080
- GRPC Rest endpoint on :8081
- Some unit and integration tests
- Cache service with Redis
- Docker compose that will release the system

## Kaspar API

#### Get a recomendation based on a specific stock and date

```http
  GET /v1/stocks/:name/*date
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. Name of the stock |
| `date` | `string` | **Optional**. Date of the stock that you want the recomendation. If not provided assumes the current date. |

#### Health

```http
  GET /health
```

Returns 200 if the system and its dependencies are up and running. If anything is wrong returns 500.

#### Ping

```http
  GET /ping
```

Returns 200 if the system is up. 

## TODO
- Prometheous metrics
- Grafana configuration
- E2E tests on docker-compose
- Consider connection multiplexing golang (Tried cmux. Do this instead: https://github.com/gin-gonic/gin/issues/2501)