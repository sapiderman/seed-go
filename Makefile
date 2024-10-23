# seed-go makefile

CURRENT_PATH ?= $(shell pwd)
IMAGE_NAME ?= mock-go-img

.PHONY: all test clean build docker

build: test-static
	# GOOS=darwin GOARCH=amd64 go build -a -o $(IMAGE_NAME) cmd/Main.go  # for mac
	go build -a -ldflags '-extldflags "-static"' -o $(IMAGE_NAME) cmd/Main.go # for linux

build-lint:
	go fmt ./...

clean:
	go clean
	rm -f $(IMAGE_NAME)

test-static: build-lint
	go vet ./...
	staticcheck ./...
	govulncheck -show verbose ./...

test-short: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out -short

test: lint
	go test ./... -v -race -covermode=atomic -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out

docker:
	docker build -t $(IMAGE_NAME) -f .docker/Dockerfile .

docker-run:
	docker run -p 7000:7000 -d $(IMAGE_NAME)
