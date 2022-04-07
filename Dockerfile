FROM golang:1.18-alpine3.15 AS builder

RUN apk update && \
    apk add --no-cache git

WORKDIR /app
COPY . /app

RUN go mod download -x && go mod verify

RUN CGO_ENABLED=0 go build -o /app/build/simple-go

FROM alpine:3.15

COPY --from=builder /app/build/simple-go /go/bin/simple-go

ENTRYPOINT ["/go/bin/simple-go"]
