# build
FROM golang:1.17-alpine as buildenv

WORKDIR /app

COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -tags=nomsgpack -ldflags="-s -w" -v -o /url-shortener

# deploy
FROM alpine

WORKDIR /

RUN mkdir /config

COPY --from=buildenv /url-shortener /url-shortener
COPY --from=buildenv /app/config /config

ARG SERVER_PORT=8080

EXPOSE ${SERVER_PORT}

CMD ["./url-shortener", "-port", "${SERVER_PORT}", "-db", "redis"]
