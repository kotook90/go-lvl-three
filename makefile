.PHONY: test
test:
	go test ./...

.PHONY: lint
lint: test
	golangci-lint run ./...

.PHONY: build
build: lint
	go build main.go

