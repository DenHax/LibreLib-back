OS := $(shell uname)

BASH_AUTO := ./scripts/autostart.sh
POWERSHELL_AUTO := ./scripts/autostart.ps1

all: compose migrate-up migrate-drop migrate-down psql-start in-psql auto-start

compose-run:
	. ./scripts/deploy/docker-compose_start.sh

compose-down:
	. ./scripts/deploy/docker_compose_down.sh

compose-env:
	. ./scripts/gen/env-compose.sh

compose-autostart:
	@$(MAKE) compose-env
	@$(MAKE) compose-run

migrate-up:
	. ./scripts/utils/migrate-up.sh

migrate-down:
	. ./scripts/utils/migrate-down.sh

migrate-drop:
	. ./scripts/utils/migrate-drop.sh

psql-start:
	. ./scripts/deploy/psql_start.sh

look-psql:
	. ./scripts/utils/lookup-in-psql.sh

run-serv:
	. ./scripts/deploy/run-serv.sh


.PHONY: auto-start compose
