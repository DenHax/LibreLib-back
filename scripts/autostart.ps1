#!/usr/bin/env pwsh

@"
$env:POSTGRES_PASSWORD = "p4ssw0rd"
$env:POSTGRES_PORT = "5432"
$env:POSTGRES_HOST = "postgres"
$env:DB_USER = "librelib-admin"
$env:DB_NAME = "librelib"
$env:DATABASE_URL = "postgres://$($env:DB_USER):$($env:POSTGRES_PASSWORD)@$($env:POSTGRES_HOST):$($env:POSTGRES_PORT)/$($env:DB_NAME)?sslmode=disable"
"@ | Set-Content -Path ..\.env.ps1

. ..\.env.ps1

Write-Host "Environment activation: succeeded"
