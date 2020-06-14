# seed-go makefile

IMAGE_NAME ?= seed-go-img

.PHONY: all test clean build docker

clean:
	go clean
	rm -f $(IMAGE_NAME)

build: 
	go build -a -o $(IMAGE_NAME) cmd/Main.go

lint: build
	golint -set_exit_status ./...

test-short: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out -short

test: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out

docker:
	docker build -t seed-go-img-dev -f Dockerfile .

	


