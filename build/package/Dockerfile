FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["../../go.mod", "../../go.sum", "./"]

RUN go mod download

COPY ../../cmd ./cmd
# COPY ../../internal ./internal

RUN go build -o ./bin/librelib ./cmd/librelib/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/librelib / 
# COPY ../configs/config.yaml /config.yaml

CMD ["/librelib"]
