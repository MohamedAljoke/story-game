.PHONY: run run-term build build-term build-all vet clean

run:
	go run ./cmd/game/

run-term:
	go run ./cmd/game-raylib/

build:
	go build -o bin/dream-walker ./cmd/game/

build-term:
	go build -o bin/dream-walker-term ./cmd/game-raylib/

build-all: build build-term

vet:
	go vet ./...

clean:
	rm -rf bin/
