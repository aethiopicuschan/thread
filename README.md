# aethiopicuschan/thread

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)

Golang library to get 5ch threads.

## Getting Started

```sh
go get github.com/aethiopicuschan/thread
```

## Usage

```go
package main

import "github.com/aethiopicuschan/thread"

func main() {
  threads, err := thread.Get("https://greta.5ch.net/poverty/")
  ...
}
```

## Running Tests

```sh
go test
```