compose:
	. ./scripts/docker-compose_start.sh

migrate-up:
	. ./scripts/migrate-up.sh

migrate-down:
	. ./scripts/migrate-down.sh

migrate-drop:
	. ./scripts/migrate-drop.sh

psql-start:
	. ./scripts/psql_start.sh

in-psql:
	. ./scripts/lookup-in-psql.sh
