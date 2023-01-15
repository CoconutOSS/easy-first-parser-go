COVERAGE = coverage.out

all: build

.PHONY: deps
deps:
	go get github.com/mattn/goveralls
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: bindata
bindata