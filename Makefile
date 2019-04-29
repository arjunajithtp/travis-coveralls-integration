LIB_PKGS=$(shell ls -d */ grep -v scripts)

default: install build

.PHONY: install
install:
	@GO111MODULE=on go install ./$(LIB_PKGS[@])...

.PHONY: tests 
tests: 
	@GO111MODULE=on go test ./$(LIB_PKGS[@])... -v

.PHONY: lintchecks 
lintchecks: 
	@GO111MODULE=on go get -v golang.org/x/lint/golint
	@GO111MODULE=on golint ./$(LIB_PKGS[@])...
