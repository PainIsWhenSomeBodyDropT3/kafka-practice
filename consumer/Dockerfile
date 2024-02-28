FROM golang:1.21-alpine as build

WORKDIR /app

COPY . .

RUN apk add librdkafka-dev pkgconf && \
  apk add build-base git bash alpine-sdk && \
  go get -tags musl 

RUN CGO_ENABLED=1 go build -tags musl -o consumer . 

FROM alpine:3 

WORKDIR /app

COPY --from=build /app/consumer .

ENTRYPOINT ["./consumer"]


