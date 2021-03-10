# Go parameters
APP_NAME=swapi
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_DIR=bin
BINARY_NAME=${APP_NAME}
BINARY_PATH=$(BINARY_DIR)/$(BINARY_NAME)
IMAGE=${APP_NAME}

all: up dep run

run:
	$(GORUN) api/main.go

swapi:
	$(GORUN) internal/swapi.go

build: dep
	GOOS=linux GOARCH=amd64 $(GOBUILD) api/main.go -o ../$(BINARY_PATH) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)
	rm -f $(BINARY_UNIX)

dep:
	go mod vendor

deps-install:
	$(GOCMD) install ./...

test: test-cover view-cover

test-unit: up dep
	$(GOTEST) -v ./...

test-cover: up dep
	$(GOTEST) ./... -covermode=count -coverprofile=count.out

view-cover:
	$(GOCMD) tool cover -html=count.out

up:
	docker-compose up -d

down:
	docker-compose down