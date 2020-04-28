MODULE_NAME:=post

.PHONY: build
build:

	go build -o cmd/${MODULE_NAME}-web cmd/main.go cmd/plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t ${MODULE_NAME}-web:latest
