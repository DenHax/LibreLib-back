FROM golang:1.21-alpine AS builder

# WORKDIR /usr/local/src

WORKDIR /app

RUN apk add --no-cache bash git make gcc gettext envsubst musl-dev 

COPY ../../go.mod .

RUN go mod download

COPY ../../cmd ./cmd

RUN go build -o ./bin/librelib ./cmd/librelib/main.go

FROM alpine AS runner

# COPY --from=builder /usr/local/src/bin/librelib /

COPY --from=builder /app/bin/librelib /

RUN apk add --no-cache ca-certificates

# COPY ../../configs/config.yaml /config.yaml

CMD ["/librelib"]
