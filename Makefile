APP?=vertigo

PORT?=8000
LOG_LEVEL?=debug
DATABASE?=postgres://pavelmurkin:Ghjcnj1987@localhost:5432/advertservice?sslmode=disable

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
BRANCH:=$(shell git branch | sed -n -e 's/^\* \(.*\)/\1/p')

PROJECT?=github.com/kot13/vertigo

GOOS?=linux
GOARCH?=amd64

SPEC?=./spec.yml

LDFLAGS?=-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME} -X ${PROJECT}/version.Branch=${BRANCH}"

clean:
	rm -f ${APP}

dep:
	gvt restore
	
gen:
	# generate code under develop

migrate:
	DATABASE=${DATABASE} go run ./migrations/*.go

compile: 
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -o ${APP}

build: clean gen dep compile
		
container: build
	docker build -t $(APP):$(RELEASE) .

run: container
	docker stop ${APP} || true && docker rm ${APP} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" -e "LOG_LEVEL=${LOG_LEVEL}" -e "DATABASE=${DATABASE}" \
		$(APP):$(RELEASE)

rund: container
	docker stop ${APP} || true && docker rm ${APP} || true
	docker run -d --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" -e "LOG_LEVEL=${LOG_LEVEL}" -e "DATABASE=${DATABASE}" \
		$(APP):$(RELEASE)

runLocal: clean gen
	go build ${LDFLAGS} -o ${APP}
	PORT=${PORT} LOG_LEVEL=${LOG_LEVEL} DATABASE=${DATABASE} ./${APP}

test: rund
	PORT=${PORT} go test -v -race ./e2e/...
