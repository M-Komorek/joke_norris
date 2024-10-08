build:
	go build -o bin/joke_norris

run: build
	./bin/joke_norris

test:
	go test -v ./...