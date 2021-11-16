# go-execution-client

[![Tag](https://img.shields.io/github/tag/attestantio/go-execution-client.svg)](https://github.com/attestantio/go-execution-client/releases/)
[![License](https://img.shields.io/github/license/attestantio/go-execution-client.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/attestantio/go-execution-client?status.svg)](https://godoc.org/github.com/attestantio/go-execution-client)
![Lint](https://github.com/attestantio/go-execution-client/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/attestantio/go-execution-client)](https://goreportcard.com/report/github.com/attestantio/go-execution-client)

Go library providing an abstraction to Ethereum execution nodes.  Its external API follows the official [Ethereum consensus APIs](https://github.com/ethereum/consensus-APIs) specification.

This library is under development; expect APIs and data structures to change until it reaches version 1.0.  In addition, clients' implementations of both their own and the standard API are themselves under development so implementation of the the full API can be incomplete.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-execution-client` is a standard Go module which can be installed with:

```sh
go get github.com/attestantio/go-execution-client
```

## Support

`go-execution-client` supports execution nodes that comply with the standard execution node API.

## Usage

Please read the [Go documentation for this library](https://godoc.org/github.com/attestantio/go-execution-client) for interface information.

## Example

Below is a complete annotated example to access an execution node.

```
package main

import (
    "context"
    "fmt"

    execclient "github.com/attestantio/go-execution-client"
    "github.com/attestantio/go-execution-client/jsonrpc"
    "github.com/rs/zerolog"
)

func main() {
    // Provide a cancellable context to the creation function.
    ctx, cancel := context.WithCancel(context.Background())
    client, err := jsonrpc.New(ctx,
        // WithAddress supplies the address of the execution node, as a URL.
        jsonrpc.WithAddress("http://localhost:8545/"),
        // LogLevel supplies the level of logging to carry out.
        jsonrpc.WithLogLevel(zerolog.WarnLevel),
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("Connected to %s\n", client.Name())

    // Client functions have their own interfaces.  Not all functions are
    // supported by all clients, so checks should be made for each function when
    // casting the service to the relevant interface.
    if provider, isProvider := client.(execclient.ChainHeightProvider); isProvider {
        chainHeight, err := provider.ChainHeight(ctx)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Chain height is %v\n", chainHeight)
    }

    // Cancelling the context passed to New() frees up resources held by the
    // client, closes connections, clears handlers, etc.
    cancel()
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/attestantio/go-execution-client/issues).

## License

[Apache-2.0](LICENSE) © 2021 Attestant Limited
