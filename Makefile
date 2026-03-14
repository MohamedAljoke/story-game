.PHONY: run run-raylib build build-raylib build-all vet clean

UNAME := $(shell uname -s)
ifeq ($(UNAME),Linux)
	RAYLIB_TAGS := -tags x11
endif

run:
	go run ./cmd/game/

run-raylib:
	go run $(RAYLIB_TAGS) ./cmd/game-raylib/

build:
	go build -o bin/dream-walker ./cmd/game/

build-raylib:
	go build $(RAYLIB_TAGS) -o bin/dream-walker-raylib ./cmd/game-raylib/

build-all: build build-raylib

vet:
	go vet ./...

clean:
	rm -rf bin/
