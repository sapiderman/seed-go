# Golang Seed

![Go](https://github.com/sapiderman/seed-go/workflows/Go/badge.svg)  

Small starter project to play around with github actions and various pipelines.  

## Features

- gorilla mux  
- contexts  
- logrus  

todo:  

- circuit breaker  
- swagger docs  
- profiler  

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

## some docs  

[license](./LICENSE)  
[code of conduct](./code_of_conduct.md)  

## integrations  

Using some free ci tools  
[github integations](www.github.com/features/actions)  

[azure](dev.azure.com)  
[![Build Status](https://dev.azure.com/sapiderman/seed-go/_apis/build/status/sapiderman.seed-go?branchName=master)](https://dev.azure.com/sapiderman/seed-go/_build/latest?definitionId=1&branchName=master)  
[circleci](circleci.com)  
[![Build Status](https://travis-ci.com/sapiderman/seed-go.svg?branch=master)](https://travis-ci.com/sapiderman/seed-go)  
[travic-ci.com](https://travis-ci.com)  
[![Build Status](https://travis-ci.com/sapiderman/seed-go.svg?branch=master)](https://travis-ci.com/sapiderman/seed-go)

using some free code quality scanners.  
everyone should be using these, no excuses.  

[go report card](goreportcard.com)  
[![Go Report Card](https://goreportcard.com/badge/github.com/sapiderman/seed-go)](https://goreportcard.com/report/github.com/sapiderman/seed-go)

[codefactor.io](https://www.codefactor.io)  
[![CodeFactor](https://www.codefactor.io/repository/github/sapiderman/seed-go/badge)](https://www.codefactor.io/repository/github/sapiderman/seed-go)  

[codeclimate.com](https://www.codeclimate.com)  
[![Maintainability](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/maintainability)](https://codeclimate.com/github/codeclimate/codeclimate/maintainability)  
[![Test Coverage](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/test_coverage)](https://codeclimate.com/github/codeclimate/codeclimate/test_coverage)

---

## project structure

TBD.

---  

fork. clone. contribute and share!  
