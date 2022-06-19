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

.PHONY: run-k8s
run-k8s:
	chmod +x ./scripts/k8s/RunKubernetes.sh
	./scripts/k8s/RunKubernetes.sh
prerequisites: run-k8s
target: prerequisites

.PHONY: stop-k8s
stop-k8s:
	chmod +x ./scripts/k8s/StopKubernetes.sh
	./scripts/k8s/StopKubernetes.sh
prerequisites: stop-k8s
target: prerequisites
