build: clean
	npm run build
	go build

run:
	./goscript

clean:
	rm ./goscript