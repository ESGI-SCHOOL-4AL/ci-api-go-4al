FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build

FROM alpine:3.9

WORKDIR /app

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY --from=builder /app/ci-api-go-4al ./ci-api-go-4al
COPY resources /app/resources
COPY config.production.json ./config.production.json

EXPOSE 80
ENV GOYAVE_ENV=production

CMD dockerize -wait tcp://mariadb:3306 ./ci-api-go-4al