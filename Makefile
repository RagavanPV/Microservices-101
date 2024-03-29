THIS_FILE := $(lastword $(MAKEFILE_LIST))
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
.PHONY: help build up start down destroy stop restart logs ps start-dependencies destroy-dependencies
help:
		make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
build:
		docker-compose -f docker-compose.yml build $(c)
up:
		docker-compose -f docker-compose.yml up -d $(c)
start:
		docker-compose -f docker-compose.yml start $(c)
down:
		docker-compose -f docker-compose.yml down $(c)
destroy:
		docker-compose -f docker-compose.yml down -v $(c)
stop:
		docker-compose -f docker-compose.yml stop $(c)
restart:
		docker-compose -f docker-compose.yml stop $(c)
		docker-compose -f docker-compose.yml up -d $(c)
logs:
		docker-compose -f docker-compose.yml logs --tail=100 -f $(c)
ps:
		docker-compose -f docker-compose.yml ps
start-dependencies:
		cd docker-compose/databases/ && docker-compose -f docker-compose.yml up -d
		cd docker-compose/elastic-stack/ && docker-compose -f docker-compose.yml up -d
		cd docker-compose/real-time-migration/ && docker-compose -f docker-compose.yml up -d
destroy-dependencies:
		cd docker-compose/databases/ && docker-compose -f docker-compose.yml down -v && rm -rf data/
		cd docker-compose/elastic-stack/ && docker-compose -f docker-compose.yml down -v && rm -rf data/ certs/
		cd docker-compose/real-time-migration/ && docker-compose -f docker-compose.yml down -v && rm -rf data/