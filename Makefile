build: clean
	npm run build
	go build

run: build
	./goscript

.PHONY: clean
clean:
	rm -f ./goscript

all: run