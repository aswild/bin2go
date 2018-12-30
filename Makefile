.NOTPARALLEL:

BINNAME = bin2go

all: build

build:
	go build -v -o $(BINNAME) ./cmd

clean:
	rm -f $(BINNAME)

goclean: clean
	go clean -cache

.PHONY: all build clean goclean
