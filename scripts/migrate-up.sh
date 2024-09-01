#! /usr/bin/env bash
migrate -path ./third_party/postgres/schema -database 'postgres://postgres:$POSTGRES_PASSWORD@POSTGRES_HOST:POSTGRES_PORT/postgres?sslmode=disable' up
