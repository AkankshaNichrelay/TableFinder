APP = tablefinder
DOCKER_IMAGE = tablefinder-build
PKG = ...

deps:
	@go mod tidy -compat=1.17

wire:
	@go get -d github.com/google/wire/cmd/wire
	@cd cmd/$(APP) && wire

test:
	@go test -race ./$(PKG)

build:
	@echo "building app"
	@go build -o ./bin/$(APP) ./cmd/$(APP)

run:
	@bin/$(APP)

docker-up:
	@echo "running docker_up"
	@docker-compose up -d

docker-down:
	@echo "running docker_down"
	@docker-compose down

docker-build:
	@echo "building docker image"
	@docker build -t $(DOCKER_IMAGE) .

docker-cleanup:
	@echo "removing docker image"
	@docker rm $(DOCKER_IMAGE)-app || true

  docker-tests: docker-build docker-up
	@echo "running docker tests"
	@docker run -p 3000:3000 --name $(DOCKER_IMAGE)-app $(DOCKER_IMAGE) \ 
	go test -race ./$(PKG)
	@make docker-down

docker-run: docker-cleanup
	@echo "running docker"
	@docker run -p 3000:3000 --name $(DOCKER_IMAGE)-app $(DOCKER_IMAGE)