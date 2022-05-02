APP = tablefinder

deps:
	@go mod download
	@go mod tidy -go=1.16 && go mod tidy -go=1.17

wire:
	@go get -d github.com/google/wire/cmd/wire
	@cd cmd/$(APP) && wire

build:
	@echo "building app"
	@go build .\cmd\$(APP)

run:
	@.\$(APP).exe
