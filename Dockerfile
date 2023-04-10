FROM golang:1.19-alpine AS builder

RUN mkdir /app

COPY . /app
WORKDIR /app

RUN CGO_ENABLE=0 go build -o recipet ./cmd/app

RUN chmod +x /app/recipet

FROM alpine:latest

RUN mkdir app

#COPY --from=builder /app/tgBot /app/
#COPY --from=0 /app/configs/main.yml ./configs/
#COPY --from=0 /app/configs/app.env ./configs/
#
#CMD ["./app/tgBot"]