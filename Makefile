.DEFAULT_GOAL := buildandrun

services:
	docker-compose --env-file .env up -d

test: services
	go test ./... -cover

build: clean
	go build main.go

run: services
	./main

buildandrun: build run

prep:
	echo "Preparing environment..."

clean:
	rm -f main

docs:
	go doc --all 

docker-build: gen-version
	# Build the docker image
	docker build -t telegramsvc .

	# Tag the image
	docker tag telegramsvc katesclau/telegramsvc:${VERSION}

docker-push: docker-build
	# Push the image to docker hub
	docker push katesclau/telegramsvc:${VERSION}

local-deploy:
	minikube start
	kubectl apply -f ./infra/deployment.yml

gen-version:
	VERSION=$(shell cat version)
