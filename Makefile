APP = tablefinder

deps:
	@go mod download
	@go mod tidy

wire:
	@go get -d github.com/google/wire/cmd/wire
	@cd cmd/$(APP) && wire

build:
	@echo "building app"
	@go build .\cmd\$(APP)
