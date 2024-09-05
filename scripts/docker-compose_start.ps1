#!/usr/bin/env pwsh

docker-compose -f ./deployments/compose.yaml up --force-recreate
