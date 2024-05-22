FROM golang:1.17-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./server/go.mod ./server/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY ../server/ .

# Build the Go app
RUN go build -o main .

# Start a new stage to ignore unneeded go source files
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY server/assets ./assets

COPY client/client ./client

# Command to run the executable
CMD ["./main"]
