.PHONY: test ctest covdir coverage docs linter qtest clean
APP_VERSION:=$(shell cat VERSION | head -1)
GIT_COMMIT:=$(shell git describe --dirty --always)
GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD -- | head -1)
BUILD_USER:=$(shell whoami)
BUILD_DATE:=$(shell date +"%Y-%m-%d")
BINARY:="go-cisco-nx-api-client"
VERBOSE:=-v
ifdef TEST
	TEST:="-run ${TEST}"
endif

all:
	@echo "Version: $(APP_VERSION), Branch: $(GIT_BRANCH), Revision: $(GIT_COMMIT)"
	@echo "Build on $(BUILD_DATE) by $(BUILD_USER)"
	@mkdir -p bin/
	@CGO_ENABLED=0 go build -o bin/$(BINARY) $(VERBOSE) \
		-ldflags="-w -s \
		-X main.appName=$(BINARY) \
		-X main.appVersion=$(APP_VERSION) \
		-X main.gitBranch=$(GIT_BRANCH) \
		-X main.gitCommit=$(GIT_COMMIT) \
		-X main.buildUser=$(BUILD_USER) \
		-X main.buildDate=$(BUILD_DATE)" \
		-gcflags="all=-trimpath=$(GOPATH)/src" \
		-asmflags="all=-trimpath $(GOPATH)/src" cmd/client/*
	@echo "Done!"

linter:
	@golint pkg/client/*.go
	@echo "PASS: golint"

test: covdir linter
	@go test $(VERBOSE) -coverprofile=.coverage/coverage.out ./pkg/client/*.go

ctest: covdir linter
	@richgo version || go get -u github.com/kyoh86/richgo
	@time richgo test $(VERBOSE) "${TEST}" -coverprofile=.coverage/coverage.out ./pkg/client/*.go

covdir:
	@mkdir -p .coverage

coverage:
	@go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html

docs:
	@mkdir -p .doc
	@godoc -html github.com/greenpau/gonxapiclient > .doc/index.html
	@echo "Run to serve docs:"
	@echo "    godoc -goroot .doc/ -html -http \":5000\""

clean:
	@rm -rf .doc
	@rm -rf .coverage
	@rm -rf bin/

qtest:
	@#go test -v -run TestParseShowVersionJsonOutput ./pkg/client/*.go
	@#go test -v -run TestParseShowVlanJsonOutput ./pkg/client/*.go
	@#go test -v -run TestParseShowInterfaceJsonOutput ./pkg/client/*.go
	@#go test -v -run TestParseShowInterfaceEthernet ./pkg/client/*.go
	@#go test -v -run TestParseShowInterfaceSvi ./pkg/client/*.go
	@#go test -v -run TestParseShowInterfaceMgmt ./pkg/client/*.go
	@go test -v -run TestGetSystemInfo ./pkg/client/*.go
