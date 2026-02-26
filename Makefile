ifneq (,$(wildcard ./.env))
	include .env
	export
endif

export USERNAME=root
export PASSWORD=pass

export DB_PORT=27017
export MONGO_DB_NAME=GoDB
export MONGO_URI=mongodb://$(USERNAME):$(PASSWORD)@mongo:$(DB_PORT)/?authSource=admin&w=majority

.PHONY: up down ps

up:
	@docker compose -f ./docker-compose.yml up --build -d
	
down:
	@docker compose -f ./docker-compose.yml down -v

ps:
	@docker compose -f ./docker-compose.yml ps -a