# Go parameters
BINARY_NAME=k8s-annotator
DOCKER_IMAGE=immnan/k8s-annotator
TAG=latest

.PHONY: all build docker push clean

all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME) main.go

# Build the Docker image
docker:
	docker build -t $(DOCKER_IMAGE) .
	docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE):$(TAG)
	docker push $(DOCKER_IMAGE):$(TAG)

# Clean up build artifacts
clean:
	rm -f $(BINARY_NAME)

