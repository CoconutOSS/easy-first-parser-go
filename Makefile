COVERAGE = coverage.out

all: build

.PHONY: deps
deps:
	go get github.com/m