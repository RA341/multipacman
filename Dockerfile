FROM node:18-alpine AS minifier

WORKDIR /app

COPY . .

RUN npm install uglify-js -g

RUN find . -name '*.js' -exec sh -c 'uglifyjs "$1" --compress --mangle --output "$1"' _ {} \;

FROM golang:1.22-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk update && \
    apk add --no-cache gcc musl-dev

COPY --from=minifier /app /app

RUN go get

RUN go build -ldflags "-s -w" -o app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

CMD ["./app"]
