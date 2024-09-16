# Проверка на существование файла .env
if (Test-Path -Path ".\.env") {
    Write-Host "Error: .env file already exists"
} else {
    # Создание .env файла
    @"
POSTGRES_PASSWORD=p4ssw0rd
POSTGRES_PORT=5438
POSTGRES_HOST=127.0.0.1
DB_USER=librelib-admin
DB_NAME=librelib
DATABASE_URL=postgres://${env:DB_USER}:${env:POSTGRES_PASSWORD}@${env:POSTGRES_HOST}:${env:POSTGRES_PORT}/${env:DB_NAME}?sslmode=disable
"@ | Set-Content -Path ".\.env"
    Write-Host ".env file created successfully"
}

# Загрузка переменных окружения
if (Test-Path -Path ".\.env") {
    Get-Content -Path ".\.env" | ForEach-Object {
        if ($_ -and $_ -notmatch '^#') {
            $name, $value = $_ -split '=', 2
            [System.Environment]::SetEnvironmentVariable($name.Trim(), $value.Trim(), [System.EnvironmentVariableTarget]::Process)
        }
    }
    Write-Host "Environment activation: succeeded"
} else {
    Write-Host "Error: .env file not found"
}
