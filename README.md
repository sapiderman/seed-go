# Golang Seed

Small repo to play around with github actions and pipelines.

## endpoints

GET /hello  
GET /health  

## build

go build ./...  

## test

go test ./... -v -covermode=count -coverprofile=coverage.out  

## generate-binary

go build -a -o seed-go-img cmd/Main.go

## create docker

make docker


