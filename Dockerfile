# We have multi-stage docker build
# The first one is for building image to build the binary file (will be deleted after second second stage complete)
# The second one is for building final image with smaller base image with binary file only

# ==========================================
# 1st Stage
# ==========================================
FROM golang:1.16 AS builder
LABEL maintainer="76b900@gmail.com"

## Set the working directory
WORKDIR /app

## Copy source
COPY . .

## Compile
RUN make build

# ==========================================
# 2nd Stage
# ==========================================
FROM alpine:latest

ENV APP_NAME=go-boiler

WORKDIR /app

## Add ssl cert
RUN apk add --update --no-cache ca-certificates

## Copy binary file from 1st stage
COPY --from=builder /app/bin/* ./

## Copy migration files
COPY ./database ./database

CMD ["./go-boiler", "server"]
