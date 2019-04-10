
GOLIST=govendor list -no-status +local

default: check

check: lint vet test

build:
	go build

lint:
	golint -set_exit_status $$($(GOLIST))
	golint -set_exit_status .

vet:
	go vet . $$($(GOLIST))

test:
	go test -cover -v . $$($(GOLIST))

deps:
	go get -u github.com/kardianos/govendor
	go get -u golang.org/x/lint/golint
