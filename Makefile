include envs/local_region.env
include envs/local_region.secret

APP_NAME 	:= band-app
TAG 		:= $$(git log -1 --pretty=%h)
IMG 		:= $(APP_NAME):$(TAG)

# DOCKER TASKS
# Build
build:
	docker build -t $(IMG) .

# Run
run:
	TODO

# GO TASKS
# Build
gobuild:
	go build ./cmd/band-app

gobuildmongo:
	go build ./cmd/mongo-testing

# Run
gorunmongo:
	export MONGO_USERNAME=${MONGO_USERNAME}; \
	export MONGO_PASSWORD=${MONGO_PASSWORD}; \
	export MONGO_HOST=${MONGO_HOST}; \
	export MONGO_PORT=${MONGO_PORT}; \
	export MONGO_SCHEME=${MONGO_SCHEME}; \
	go run ./cmd/mongo-testing/main.go