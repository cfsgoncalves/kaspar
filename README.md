# kaspar
A small project with the objective to explore GRPC and a cache integration with REST API and GRPC.
The project consists on an endpoint that will refresh the most recent 30 stocks with a set of atributes from reddit. The project has a cache-on-the-side pattern and is available using REST or grpc.
Is also building in go using a clean-ish architecture.

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