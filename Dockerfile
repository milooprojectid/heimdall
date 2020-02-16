FROM golang:1.13-alpine AS build-env

WORKDIR /go/src/app
COPY . .
RUN bin/setup

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/app/main /app/
ENTRYPOINT ./main
