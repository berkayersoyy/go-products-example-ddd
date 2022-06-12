.PHONY: build
build:
	$(info Building)
	go build cmd/main.go

.PHONY: run
run:
	$(info Running)
	go run cmd/main.go

.PHONY: test
test:
	$(info Test started)
	go test ./...

.PHONY: test-with-coverage
test-with-coverage:
	$(info Test with coverage started)
	go test -race -parallel 10 ./... -coverprofile=./coverage.out

.PHONY: lint
lint:
	$(info Lint started)
	golangci-lint run
