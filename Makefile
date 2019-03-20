.DEFAULT_GOAL := default
.PHONY: default
default:
	docker-compose build
	docker-compose up
