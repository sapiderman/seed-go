# Golang Seed

![Go](https://github.com/sapiderman/seed-go/workflows/Go/badge.svg)  

Small starter project to play around with golang features, github actions and various pipelines.  

## Features

| Features          | Description                                                 |
| :-----------------| :---------------------------------------------------------: |  
| gorilla mux       | [routing using gorilla mux](https://github.com/gorilla/mux) |  
| contexts          | [contexts](https://golang.org/pkg/context/)                 |
| logrus            | [logging library](https://github.com/sirupsen/logrus)        |
| monitoring        | [elastic's application performance monitoring](https://www.elastic.co/guide/en/apm/agent/go/1.x/getting-started.html) |  
| health check | [IETF RFC Health Ceck via nelkinda's health-go](https://tools.ietf.org/id/draft-inadarei-api-health-check-04.html)|
| swagger api       | [API documentations using swagger swagger.io](https://swagger.io/specification/) |  
| | |  

todo:  

- circuit breaker  
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

## project structure


├── api  
├── cmd  
├── docs  
└── internal   
   
| name          | descriptions                                       |  
| :------------ | :------------------------------------------------: |
| api           | swagger api, any protocol and schema file are here |  
| cmd           | main application lives here                        |  
| docs          | design and user documentations collected here      |  
| internal      | interal application and libraries are here         |  
  
Further information see:  
1. [Golang Project Structure](https://tutorialedge.net/golang/go-project-structure-best-practices)  
2. [Golang standard project layout ](https://github.com/golang-standards/project-layout)  

## integrations  

Using some free ci tools and code quality scanners.  There's some really great tools for public projects, trying them here.

| Service           | Status                                               |
| :-------------    | :-----------------------------------:                 |
| [github integations](www.github.com/features/actions)     |               |
| [github codeQL](![CodeQL](https://github.com/sapiderman/seed-go/workflows/CodeQL/badge.svg)) | ![CodeQL](https://github.com/sapiderman/seed-go/workflows/CodeQL/badge.svg) |
| [azure](dev.azure.com) | [![Build Status](https://dev.azure.com/sapiderman/seed-go/_apis/build/status/sapiderman.seed-go?branchName=master)](https://dev.azure.com/sapiderman/seed-go/_build/latest?definitionId=1&branchName=master)               |
| [circleci](circleci.com) | [![Build Status](https://travis-ci.com/sapiderman/seed-go.svg?branch=master)](https://travis-ci.com/sapiderman/seed-go) |
| [travic-ci.com](https://travis-ci.com) |[![Build Status](https://travis-ci.com/sapiderman/seed-go.svg?branch=master)](https://travis-ci.com/sapiderman/seed-go)             |
| [appveyor.com](https://appveyor.com) | [![Build status](https://ci.appveyor.com/api/projects/status/dd8phuty1k4n4v23/branch/master?svg=true)](https://ci.appveyor.com/project/Budhi/seed-go/branch/master) |
| [goreportcard.com](https://goreportcard.com)  | [![Go Report Card](https://goreportcard.com/badge/github.com/sapiderman/seed-go)](https://goreportcard.com/report/github.com/sapiderman/seed-go)             |
| [codeclimate.com](https://www.codeclimate.com) | [![Maintainability](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/maintainability)](https://codeclimate.com/github/codeclimate/codeclimate/maintainability) |
|                   | [![Test Coverage](https://api.codeclimate.com/v1/badges/a99a88d28ad37a79dbf6/test_coverage)](https://codeclimate.com/github/codeclimate/codeclimate/test_coverage)             |  
| | |  

fork. clone. contribute and share!  
