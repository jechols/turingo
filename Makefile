.PHONY: all

all:
	vgo build

lint:
	find . -name "*.go" | xargs golint
