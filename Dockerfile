# Flutter build
FROM ghcr.io/cirruslabs/flutter:stable AS flutter_builder

RUN flutter config --enable-web --no-cli-animations && flutter doctor

WORKDIR /app/

COPY ./frontend/pubspec.* .

RUN flutter pub get

COPY ./frontend .

RUN flutter build web

FROM node:23-alpine AS node

WORKDIR /game/

COPY frontend-js/package.json .
COPY frontend-js/package-lock.json .

RUN npm i

COPY frontend-js .

RUN npm run docker

# Stage Go build
FROM golang:1.24-alpine AS go_builder

# for sqlite
ENV CGO_ENABLED=1

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /app

COPY core/go.mod .
COPY core/go.sum .
# cache deps
RUN go mod download

COPY core/ .
COPY --from=flutter_builder /app/build/web ./cmd/web
COPY --from=node /game/dist/ ./cmd/web/

# Build optimized binary without debugging symbols
RUN go build -ldflags "-s -w" -o multipacman cmd/server/main.go

# Stage: Final stage
FROM alpine:latest

WORKDIR /app/

COPY --from=go_builder /app/multipacman .

CMD ["./multipacman"]
