FROM golang:1.20.2 as builder

WORKDIR /app

# Install deps only to allow for build layer caching
COPY go.* /app/
RUN go mod download

COPY . /app
RUN go build -v -o server

# Smaller prod images
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
RUN mkdir /app/db
CMD ["/app/server"]