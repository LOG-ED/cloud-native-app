.PHONY: start-consul start build stop help

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build			when golang files are modified"
	@echo "  start     		to start the stack"
	@echo "  stop			to destroy the stack"

build:
	docker-compose build

build-no-cache:
	docker-compose build --no-cache

start-consul:
	docker-compose up -d consul

start:
	docker-compose up -d --remove-orphans

stop: 
	docker-compose down --remove-orphans