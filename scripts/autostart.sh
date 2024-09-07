#!/usr/bin/env bash

cat <<EOL >../.env
POSTGRES_PASSWORD=p4ssw0rd
POSTGRES_PORT=5432
POSTGRES_HOST=postgres
DB_USER=librelib-admin
DB_NAME=librelib
DATABASE_URL=postgres://\${DB_USER}:\${POSTGRES_PASSWORD}@\${POSTGRES_HOST}:\${POSTGRES_PORT}/\${DB_NAME}?sslmode=disable
EOL

#!/usr/bin/env bash

if [ -f ../.env ]; then
  export "$(grep -v '^#' ../.env | xargs)"
  echo "Environment activation: succeeded"
else
  echo "Error: .env file not found"
fi
