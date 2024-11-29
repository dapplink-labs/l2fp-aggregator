<!--
parent:
  order: false
-->

<div align="center">
  <h1> Manta-Relayer </h1>
</div>

<div align="center">
  <a href="https://github.com/eniac-x-labs/manta-relayer/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/eniac-x-labs/manta-relayer.svg" />
  </a>
  <a href="https://github.com/eniac-x-labs/manta-relayer/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/eniac-x-labs/manta-relayer.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/eniac-x-labs/manta-relayer">
    <img alt="GoDoc" src="https://godoc.org/github.com/eniac-x-labs/manta-relayer?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/eniac-x-labs/manta-relayer">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/eniac-x-labs/manta-relayer"/>
  </a>
</div>

manta-relayer is a decentralized system. The node first registers to the contract.
When more than 2/3 of the nodes have completed the signature verification, 
the signature will be submitted to the Ethereum contract, and then the fraud proof of withdrawal will be reduced.

**Note**: Requires [Go 1.22+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions please read the [Installation](https://github.com/eniac-x-labs/dapplink/) instructions. Once the dependencies are installed, run:

```bash
make build
```

Or check out the latest [release](https://github.com/eniac-x-labs/finality-node).

## Quick Start

* Import the environment variables in example.env

* Improve the configuration file in finality-node.yaml

* Database migration, run:
```bash
./manta-relayer manager-migrations
./manta-relayer node-migrations
```

* start the service
```bash
./manta-relayer manager
./manta-relayer node
```

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/eniac-x-labs/finality-node/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).
