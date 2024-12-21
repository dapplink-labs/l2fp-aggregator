<!--
parent:
  order: false
-->

<div align="center">
  <h1> Manta-Relayer </h1>
</div>

<div align="center">
  <a href="https://github.com/dapplink-labs/l2fp-aggregator/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/dapplink-labs/l2fp-aggregator.svg" />
  </a>
  <a href="https://github.com/dapplink-labs/l2fp-aggregator/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/dapplink-labs/l2fp-aggregator.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/dapplink-labs/l2fp-aggregator">
    <img alt="GoDoc" src="https://godoc.org/github.com/dapplink-labs/l2fp-aggregator?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/dapplink-labs/l2fp-aggregator">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/dapplink-labs/l2fp-aggregator"/>
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

* Improve the configuration file in manta-relayer.yaml

* start the service
```bash
./manta-relayer manager
./manta-relayer node
```

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/eniac-x-labs/finality-node/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).
