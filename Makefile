APP?=vertigo
PORT?=8000
LOG_LEVEL?=debug
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
PROJECT?=github.com/kot13/vertigo

clean:
	rm -f ${APP}

build: clean
	go build \
        -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
        -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
        -o ${APP}

run: build
	PORT=${PORT} LOG_LEVEL=${LOG_LEVEL} ./${APP}

test:
	go test -v -race ./...