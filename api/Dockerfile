FROM golang:alpine AS base

# Development phase
FROM base AS dev

WORKDIR /app

# Copy the database folder that containts
# the migrations into a root path
COPY ./database /database

# Install air software
# for hot reload in development purpose
RUN go install github.com/cosmtrek/air@latest

# Create a EdDSA key
RUN apk add openssl && openssl genpkey -algorithm ed25519 -out private.pem

# Espose the port
# This is mandatory for routing
EXPOSE 8081

# Execute air command
CMD ["air"]

# Build the binary phase
FROM base AS build

# Install some requiremenrs for build the application
RUN apk --no-cache add gcc g++ make git

WORKDIR /go/src/app

# Copy all the files into WORKDIR
COPY . .

# Install dependencies
RUN go mod tidy

# Build the binary
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/api ./main.go

# Install alpine version
FROM alpine:3.17

# Install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/bin

# Copy the binary
COPY --from=base /go/src/app/bin /go/bin

# Expose the port
EXPOSE 8081

ENTRYPOINT /go/bin/api




