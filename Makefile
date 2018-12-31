.NOTPARALLEL:

BINNAME = bin2go
GOFLAGS ?= -v
LDFLAGS ?= -s -w

all: build

build:
	go build $(GOFLAGS) -ldflags='$(LDFLAGS)' -o $(BINNAME) ./cmd

test:
	go test -v ./...

clean:
	rm -f $(BINNAME)

goclean: clean
	go clean -cache

.PHONY: all build test clean goclean
