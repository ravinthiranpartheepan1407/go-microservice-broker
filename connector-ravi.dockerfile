#base go image

FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o connectorRavi ./

RUN chmod +x /app/connectorRavi


# Building Tiny Image

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/connectorRavi /app

CMD [ "/app/connectorRavi" ]