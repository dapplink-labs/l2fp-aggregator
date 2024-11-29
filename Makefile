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

binding-msm:
	$(eval temp := $(shell mktemp))

	cat $(MANTA_SERVICE_MANAGER_ABI_ARTIFACT) \
	| jq -r .bytecode.object > $(temp)

	cat $(MANTA_SERVICE_MANAGER_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/manta_service_manager.go \
	--type MantaServiceManager \
	--bin $(temp)

	rm $(temp)

.PHONY: \
	 finality-node \
	 clean \
	 test \
	 lint \
	 binding-msm