# Flutter build
FROM ghcr.io/cirruslabs/flutter:stable AS flutter_builder

RUN flutter config --enable-web --no-cli-animations && flutter doctor

WORKDIR /app/

COPY ./frontend/pubspec.* .

RUN flutter pub get

COPY ./frontend .

RUN flutter build web

# Stage Go build
FROM golang:1.24-alpine AS go_builder

# for sqlite
ENV CGO_ENABLED=1

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR /app

COPY ./core .

RUN go mod tidy

COPY --from=flutter_builder /app/build/web ./web

# Build optimized binary without debugging symbols
RUN go build -o multipacman

# Stage: Final stage
FROM alpine:latest

WORKDIR /app/

COPY --from=go_builder /app/multipacman .

CMD ["./multipacman"]
