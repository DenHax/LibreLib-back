#!/usr/bin/env bash

if [ -f ./.env ]; then
  echo "Error: .env file already exists"
else
  cat <<EOL >.env
POSTGRES_PASSWORD=p4ssw0rd
POSTGRES_PORT=5432
POSTGRES_HOST=postgres
DB_USER=librelib-admin
DB_NAME=librelib
DATABASE_URL=postgres://\${DB_USER}:\${POSTGRES_PASSWORD}@\${POSTGRES_HOST}:\${POSTGRES_PORT}/\${DB_NAME}?sslmode=disable
EOL
  echo ".env file created successfully"
fi

if [ -f ./.env ]; then
  export "$(grep -v '^#' ./.env | xargs)"
  if [ $? -eq 0 ]; then
    echo "Environment activation: succeeded"
  else
    echo "Error: Failed to export environment variables"
  fi
else
  echo "Error: .env file not found"
fi
