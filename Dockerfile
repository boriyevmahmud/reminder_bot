# FROM golang:1.18.7-alpine3.16 AS builder

# RUN mkdir app

# COPY . /app

# WORKDIR /app

# RUN go build -o main cmd/main.go

# FROM alpine:3.16

# WORKDIR /app

# COPY --from=builder /app .
# ENV DOT_ENV_PATH=config/.env

# CMD ["/app/main"]

FROM golang:1.18.7-alpine3.16 AS builder

RUN mkdir app

COPY . /app

WORKDIR /app

# Remove the original build command
# RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

ENV DOT_ENV_PATH=config/.env

CMD ["./main"]

