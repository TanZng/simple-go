# Pet API

You need: 

🐹 Go v1.18+

🐋 Docker v20.10.14+

🦑 docker-compose v1.28.2+

## 📦 Used packages

| Package | Version | Type |
| --- | --- | --- |
| [gofiber/fiber](https://github.com/gofiber/fiber) | v2.31.0 | Core |
| [go-gorm/gorm](https://github.com/go-gorm/gorm) | v1.23.3 | Database |
| [stretchr/testify](https://github.com/stretchr/testify)| v1.7.0 | Test |
| [uber-go/zap](https://github.com/uber-go/zap) | v1.21.0 | Logs |
| [google/uuid](https://github.com/google/uuid) | v1.3.0 | Utils |
| [joho/godotenv](https://github.com/joho/godotenv) | v1.4.0 | Config |

## 🏃 Execute

Run using:

```bash
docker-compose up
```

or run locally using:

```bash
go mod download -x && go mod verify
go run main.go
```

> In order to run it locally is necessary to have a PostgreSQL instance running in your machine.

## ⤴ Routes

### GET `/hello-world`

```bash
curl -X GET http://127.0.0.1:8080/hello-world
```

### POST `pet/`

```bash
curl -X POST http://127.0.0.1:8080/pet \
   -H 'Content-Type: application/json' \
   -d '{"name":"Megan","kind":"Dog"}'
```

Example Output:

```json
{"id":"56a7b854-1dca-415a-9259-9432c993b363","name":"Megan","kind":"Dog","created-at":"0001-01-01T00:00:00Z","updated-at":"2022-04-07T16:53:59.461711837Z","deleted-at":null}
```

### GET `pet/:id`

```bash
curl -X GET http://127.0.0.1:8080/pet/7c7c62c67-473e-48f2-a7da-60f0c24f2b6b
```

Example Output:

```json
{"id":"c7c62c67-473e-48f2-a7da-60f0c24f2b6b","name":"Onix","kind":"Cat","created-at":"0001-01-01T00:00:00Z","updated-at":"2022-04-07T16:51:25.01315Z","deleted-at":null}  
```

## 🧪 Test

Run test suite:

```bash
go test ./...
```