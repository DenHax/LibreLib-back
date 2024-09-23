OS := $(shell uname)

BASH_AUTO := ./scripts/autostart.sh
POWERSHELL_AUTO := ./scripts/autostart.ps1

all: compose migrate-up migrate-drop migrate-down psql-start in-psql auto-start

compose:
ifeq ($(OS),Linux)
	. ./scripts/docker-compose_start.sh
else ifeq ($(OS),Darwin)
	. ./scripts/docker-compose_start.sh
else ifeq ($(OS),Windows_NT)
	@powershell -ExecutionPolicy Bypass -File ./scripts/docker-compose_start.ps1
endif

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

autostart:
ifeq ($(OS),Linux)
	. ./scripts/autostart.sh
else ifeq ($(OS),Darwin)
	. ./scripts/autostart.sh
else ifeq ($(OS),Windows_NT)
	@powershell -ExecutionPolicy Bypass -File ./scripts/autostart.ps1
endif
	@$(MAKE) compose

.PHONY: auto-start compose
