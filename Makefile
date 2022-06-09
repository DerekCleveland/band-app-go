include envs/local_region.env
include envs/local_region.secret

APP_NAME 	:= band-app-server
TAG 		:= $$(git log -1 --pretty=%h)
IMG 		:= ${APP_NAME}:${TAG}
PORT 		:= 8080
MONGO_PORT 	:= 27017

# DOCKER TASKS
# Build
build:
	docker build -t ${IMG} .

# Run
run:
	docker run --rm --env-file=./envs/local_region_docker.env --env-file=./envs/local_region.secret --link band-app-mongo -p=${PORT}:${PORT} ${IMG}

# GO TASKS
# Build
gobuild:
	go build ./cmd/band-app-server

gobuildmongo:
	go build ./cmd/mongo-testing

# Run
gorun:
	export MONGO_USERNAME=${MONGO_USERNAME}; \
	export MONGO_PASSWORD=${MONGO_PASSWORD}; \
	export MONGO_HOST=${MONGO_HOST}; \
	export MONGO_PORT=${MONGO_PORT}; \
	export MONGO_SCHEME=${MONGO_SCHEME}; \
	go run ./cmd/band-app-server/main.go

gorunmongo:
	export MONGO_USERNAME=${MONGO_USERNAME}; \
	export MONGO_PASSWORD=${MONGO_PASSWORD}; \
	export MONGO_HOST=${MONGO_HOST}; \
	export MONGO_PORT=${MONGO_PORT}; \
	export MONGO_SCHEME=${MONGO_SCHEME}; \
	go run ./cmd/mongo-testing/main.go