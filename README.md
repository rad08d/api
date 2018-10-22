# Simple Rate Calculator
This app will let you supply a JSON object of rates and availabe days along with a start time and end time as a query string and return if the resource is availabe and for what price.

## Requirements
1) Golang 1.10 or up
2) git

## Test
1) cd $GOPATH/src/github.com/rad08d/api
2) go test

## Running
1) go get github.com/rad08d/api
2) cd $GOPATH/src/github.com/rad08d/api
3) go run server.go
4) Application will be availabe at localhost:8080
5) Swagger docs are located at http://localhost:8080/swagger/index.html

## Build/Install/Run
1) cd $GOPATH/src/github.com/rad08d/api
2) go install
3) cd $GOPATH/bin/
4) ./api