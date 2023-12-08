# Makefile to build and run MongoDB container

# Docker image name and tag
IMAGE_NAME = mongodb
IMAGE_TAG = latest

# Docker container name
CONTAINER_NAME = mongodb-container

# Dockerfile name
DOCKERFILE = Dockerfile.mongo

# Build the MongoDB Docker image
build-mongo: 
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f $(DOCKERFILE) .

build: build-mongo
	go build -o short-it .

run-mongo:
	docker run -d -p 27017:27017 --name $(CONTAINER_NAME) $(IMAGE_NAME):$(IMAGE_TAG)

# Run the MongoDB container
run: run-mongo
	./short-it

# Stop and remove the MongoDB container
stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

.PHONY: build run stop
