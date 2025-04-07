# Pagination Seven in Go

This repository contains an implementation of Pagination Seven
(https://codingdojo.org/kata/PaginationSeven/) in Go. It is developed as a
learning practice, especially to get to know Go and also practice TDD.

## Setup

Either download the pre-build binaries from the latest release
[here](https://github.com/JoshuaGanter/go-pagination-seven/releases/latest) or
build from source.

### Build

Go >1.24.2 is required to build the project, get it
[here](https://go.dev/doc/install). Download this repository's content or clone
it.

Build with:

```shell
$ go build -o build/paginationseven
```

See https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies for more
options on `go build`.

## Usage

Run the binaries with:

```shell
$ ./build/paginationseven <current_page> <total_pages>
```

### Run with Go

You may also run the project without building binaries:

```shell
$ go run . <current_page> <total_pages>
```

### Run tests

Run the tests with:

```shell
$ go test ./...
```

## Inspiration

This project is inspired by this tutorial of TDD:
https://www.youtube.com/watch?v=H943sC8r7nQ (in German)
