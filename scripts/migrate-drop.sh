#!/usr/bin/env bash

migrate -path ./third_party/postgres/schema -database "$POSTGRES_URL" drop
