GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"
MANTA_SERVICE_MANAGER_ABI_ARTIFACT := bindings/abi/MantaServiceManager.json

build:
	env GO111MODULE=on go build -o manta-relayer ./cmd

clean:
	rm manta-relayer

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	 finality-node \
	 clean \
	 test \
	 lint \