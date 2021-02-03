#!make
.DEFAULT_GOAL=start

local.build:
	@docker-compose build

local.start:
	@docker-compose up -d

local.down:
	@docker-compose down

logs:
	@docker logs -f $(shell docker-compose ps -q pelipper)

sh:
	@docker-compose exec pelipper /bin/sh

start: local.start

stop: local.down

renew: local.down local.build local.start

.PHONY:  start stop sh logs renew
