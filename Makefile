APP?=vertigo

PORT?=8000
LOG_LEVEL?=debug

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
BRANCH:=$(shell git branch | sed -n -e 's/^\* \(.*\)/\1/p')

PROJECT?=github.com/kot13/vertigo

GOOS?=linux
GOARCH?=amd64

BASE_SWAGGER_SPEC?=./swagger/base.yml
SWAGGER_SPEC?=./swagger/swagger.json

LDFLAGS?=-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME} -X ${PROJECT}/version.Branch=${BRANCH}"

clean:
	rm -f ${APP}
	rm -rf ./models
	rm -rf ./restapi
	rm -rf ./client

dep:
	dep ensure
	
gen:
	swagger flatten ${BASE_SWAGGER_SPEC} -o ${SWAGGER_SPEC}
	swagger generate server -A ${APP} -f ${SWAGGER_SPEC} --exclude-main
	swagger generate client -A ${APP} -f ${SWAGGER_SPEC}
	
compile: 
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -o ${APP}
	

build: clean gen dep compile
		
container: build
	docker build -t $(APP):$(RELEASE) .

run: container
	docker stop ${APP} || true && docker rm ${APP} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

rund: container
	docker stop ${APP} || true && docker rm ${APP} || true
	docker run -d --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

test: rund
	PORT=${PORT} go test -v -race ./e2e/...
	
docs:
	swagger serve -F redoc ${SWAGGER_SPEC}