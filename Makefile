SHELL := /bin/bash

test: lint vet build

lint:
	golangci-lint run

vet:
	go vet ./...

build: clean
# 在windows环境中可以直接执行 go build
	go build -v -o ./bin/qf

clean:
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./dist

# Commands for docker images.
# ----------------------------
build-linux:
	GOOS=linux GOARCH=amd64 go build -v -o ./bin/qf .

build-docker: clean build-linux
	docker build -t quickfixgo/qf:latest .