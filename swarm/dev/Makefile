.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/shaft/go-shaft/cmd/swarm
	go build -o bin/geth     github.com/shaft/go-shaft/cmd/geth
	go build -o bin/bootnode github.com/shaft/go-shaft/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/shaft/go-shaft/swarm/...
