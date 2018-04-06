.PHONY: all lint test

all:
	vgo build

lint:
	find . -name "*.go" | xargs -l1 golint

test:
	vgo test ./...
