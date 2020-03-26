APP?=statsd-notifier
USERSPACE?=serboox
RELEASE?=0.0.1
PROJECT?=github.com/${USERSPACE}/${APP}
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS?=linux
REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

# Default target executed when no arguments are given to make.
default_target: run

run: golangci-lint tests run-app
	@echo "+ $@"

run-app: clean build
	@echo "+ $@"
	./${APP} --cfg-dir ./ --cfg-name config.yaml

golangci-lint: goimports
	@echo "+ $@"
	GO111MODULE=on golangci-lint run ./...

goimports:
	@echo "+ $@"
	goimports -d -w main.go ./src

clean:
	@echo "+ $@"
	rm -f ${APP}

kill:
	@echo "+ $@"
	lsof -i :${PORT} | grep ${APP} | awk '{print $$2}' | xargs kill -9

build: clean
	@echo "+ $@"
	GOFLAGS=-mod=vendor CGO_ENABLED=1 go build -v -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/src/version.Release=${RELEASE} -X ${PROJECT}/src/version.Commit=${COMMIT} -X ${PROJECT}/src/version.Repository=${REPO_INFO} -X ${PROJECT}/src/version.BuildTime=${BUILD_TIME}"

tests:
	@echo "+ $@"
	go test ./...

send-reqests:
	while :; \
 	do \
          curl -X POST http://127.0.0.1:8077?count=1 \
          	-d '{"geo":{"CityName":"New York City","ContinentCode":"NA","CountryIsoCode":"US"}}' \
            /dev/null 1>/dev/null 2>/dev/null; \
		  sleep 0.1; \
	done;




