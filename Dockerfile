FROM golang:1.19-alpine AS builder

RUN mkdir /app

COPY . /app
WORKDIR /app

RUN CGO_ENABLE=0 go build -o Recipe_test ./cmd/app

RUN chmod +x /app/Recipe_test

FROM alpine:latest

RUN mkdir app

COPY --from=builder /app/Recipe_test /app/
COPY --from=0 /app/configs/config.env ./configs/

CMD ["./app/Recipe_test"]