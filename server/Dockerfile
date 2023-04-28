# Server Dockerfile
FROM golang:1.20-alpine as build-stage

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

FROM alpine as production-stage

WORKDIR /app

COPY --from=build-stage /usr/local/bin/app /usr/local/bin/app

EXPOSE 8090
CMD ["app", "serve", "--http", "0.0.0.0:8090"]
