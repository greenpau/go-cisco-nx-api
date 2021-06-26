APP_VERSION:=$(shell cat VERSION | head -1)
GIT_COMMIT:=$(shell git describe --dirty --always)
GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD -- | head -1)
BUILD_USER:=$(shell whoami)
BUILD_DATE:=$(shell date +"%Y-%m-%d")
APP_NAME:="go-cisco-nx-api"
BINARY:="go-cisco-nx-api-client"
VERBOSE:=-v
ifdef TEST
	TEST:="-run ${TEST}"
endif

.PHONY: all
all: info build
	@echo "OK"

.PHONY: info
info:
	@echo "Version: $(APP_VERSION), Branch: $(GIT_BRANCH), Revision: $(GIT_COMMIT)"
	@echo "Build on $(BUILD_DATE) by $(BUILD_USER)"

.PHONY: build
build:
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

.PHONY: linter
linter:
	@echo "Running lint checks"
	@#golint -set_exit_status ./...
	@echo "PASS: golint"

.PHONY: test
test: covdir linter
	@go test $(VERBOSE) -coverprofile=.coverage/coverage.out ./...

.PHONY: ctest
ctest: covdir linter
	@time richgo test $(VERBOSE) $(TEST) -coverprofile=.coverage/coverage.out ./...

.PHONY: covdir
covdir:
	@echo "Creating .coverage/ directory"
	@mkdir -p .coverage

.PHONY: coverage
coverage:
	@go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html
	@go test -covermode=count -coverprofile=.coverage/coverage.out ./...
	@go tool cover -func=.coverage/coverage.out | grep -v "100.0"

.PHONY: clean
clean:
	@rm -rf .doc
	@rm -rf .coverage
	@rm -rf bin/

.PHONY: qtest
qtest:
	@echo "Perform quick tests ..."
	@time richgo test -v -run TestParseShowSystemEnvironmentJsonOutput ./pkg/client/*.go

.PHONY: dep
dep:
	@echo "Making dependencies check ..."
	@go get -u golang.org/x/lint/golint
	@go get -u golang.org/x/tools/cmd/godoc
	@go get -u github.com/kyoh86/richgo
	@go get -u github.com/greenpau/versioned/cmd/versioned
	@go get -u github.com/google/addlicense

.PHONY: license
license:
	@for f in `find ./ -type f -name '*.go'`; do addlicense -c "Paul Greenberg greenpau@outlook.com" -y 2020 $$f; done

.PHONY: release
release:
	@echo "Making release"
	@go get -u ./...
	@go mod tidy
	@go mod verify
	@if [ $(GIT_BRANCH) != "main" ]; then echo "cannot release to non-main branch $(GIT_BRANCH)" && false; fi
	@git diff-index --quiet HEAD -- || ( echo "git directory is dirty, commit changes first" && false )
	@versioned -patch
	@git add VERSION
	@git commit -m 'updated VERSION file'
	@versioned -sync cmd/client/main.go
	@echo "Patched version"
	@git add cmd/client/main.go
	@git commit -m "released v`cat VERSION | head -1`"
	@git tag -a v`cat VERSION | head -1` -m "v`cat VERSION | head -1`"
	@git push
	@git push --tags
	@echo "If necessary, run the following commands:"
	@echo "  git push --delete origin v$(APP_VERSION)"
	@echo "  git tag --delete v$(APP_VERSION)"
