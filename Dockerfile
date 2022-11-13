
FROM golang:latest AS builder

WORKDIR /app

COPY . .


RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build

FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/modlifesdrive .

RUN mkdir db

RUN mkdir files

EXPOSE 8001

ENTRYPOINT [ "/app/modlifesdrive" ]