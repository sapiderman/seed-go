![Go](https://github.com/sapiderman/seed-go/workflows/Go/badge.svg)
[![Build Status](https://dev.azure.com/sapiderman/seed-go/_apis/build/status/sapiderman.seed-go?branchName=master)](https://dev.azure.com/sapiderman/seed-go/_build/latest?definitionId=1&branchName=master)

# Golang Seed

Small starter project to play around with github actions and various pipelines.  

## Features

- gorilla mux  
- contexts  
- logrus  

## endpoints

GET /health  
GET /v1/hello  

## build

`go build ./...`  

## tests

`go test ./... -v -covermode=count -coverprofile=coverage.out`  
or  
`make test`  
`make test-coverage`  

## generate-binary

`go build -a -o seed-go-img cmd/Main.go`  
or  
`make build`  

## create docker

`make docker`  

## run docker  

`make docker-run`  

fork. clone. contribute and share!  
