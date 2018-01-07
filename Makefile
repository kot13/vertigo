APP?=vertigo

PORT?=8000
LOG_LEVEL?=debug

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

PROJECT?=github.com/kot13/vertigo

GOOS?=linux
GOARCH?=amd64

SWAGGER_SPEC?=./swagger.yml

LDFLAGS?=-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}"

clean:
	rm -f ${APP}
	
dep:
	dep ensure
	
gen: 
	swagger generate client -A ${APP} -f ${SWAGGER_SPEC}
	swagger generate server -A ${APP} -f ${SWAGGER_SPEC}

build: clean gen dep
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -o ${APP}
		
container: build
	docker build -t $(APP):$(RELEASE) .

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

test:
	go test -v -race ./...
	
docs: 
	swagger serve -F redoc ./swagger.yml