.PHONY: br

br: build
	./f1ne

build:
	./bundle.sh
	go build

run:
	./bundle.sh
	go run .
