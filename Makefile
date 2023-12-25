DOCKER_IMAGE_NAME=bankidmock
DOCKER_IMAGE_TAG=$(shell git rev-parse --short HEAD)

.PHONY: test
test:
	go test ./...
	golangci-lint run --timeout 10m

.PHONY: build
build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) -f build/Dockerfile .

.PHONY: build-dev
build-dev:
	docker build -t $(DOCKER_IMAGE_NAME):dev -f build/Dockerfile .

.PHONY: dev
dev:
	docker compose -f deployments/docker-compose.yml up
