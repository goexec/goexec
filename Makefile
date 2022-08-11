APPNAME=goexec
GIT_COMMIT:=$(shell git rev-list -1 HEAD)
BUILD_TIME:=$(shell date +%Y%m%d%H%M%S)
GIT_TAG:=$(shell git describe --tags $(git rev-list --tags --max-count=1))
DLDFLAGS:=-X main.GitCommit=${GIT_COMMIT} -X main.BuildDate=${BUILD_TIME}
LDFLAGS:=-X main.GitCommit=${GIT_COMMIT} -X main.Version=${GIT_TAG} -X main.BuildDate=${BUILD_TIME} -s -w

build:
	go build -o "$(APPNAME)" -ldflags "$(LDFLAGS)" .

dev:
	go build -o "$(APPNAME)" -ldflags "$(DLDFLAGS)" .

install:
	go install -o "$(APPNAME)" -ldflags "$(LDFLAGS)" .

test:
	go test -v ./...

clean:
	rm "$(APPNAME)"
