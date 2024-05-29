# Stage 1: Minify JavaScript files
FROM node:18-alpine AS minifier

# Set the working directory
WORKDIR /app

# Copy the JavaScript files to the working directory
COPY . .

# Install a JavaScript minifier (e.g., uglify-js)
RUN npm install uglify-js -g

# Find all JavaScript files and minify them in place
RUN find . -name '*.js' -exec sh -c 'uglifyjs "$1" --compress --mangle --output "$1"' _ {} \;

# Stage 2: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

ENV CGO_ENABLED=1

# Install gcc
RUN apk update && \
    apk add --no-cache gcc musl-dev

# Copy the source files from the previous stage
COPY --from=minifier /app /app

RUN go get

# Build the Go application
RUN go build -ldflags "-s -w" -o app

# Stage 3: Final stage - copy the Go binary to a minimal image
FROM alpine:latest

WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/app .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

# Command to run the Go application
CMD ["./app"]
