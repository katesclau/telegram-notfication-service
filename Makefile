.DEFAULT_GOAL := buildandrun

services:
	docker-compose --env-file .env up -d

build:
	go build main.go

run: services
	./main

buildandrun: build run

prep:
	echo "Preparing environment..."

clean:
	rm -f main