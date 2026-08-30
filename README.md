# go-learn-basics

Small, self-contained Go programs one concept each, numbered in the order
learned, written from recall after study.

## Examples

| # | Example | Demonstrates |
|--:|---------|--------------|
| 01 | [hello](examples/c01-hello) | Toolchain proof: package main, imports, Println |

Run any of them from the repo root:

```sh
go run ./examples/c01-hello
```

The only prerequisite is a [Go install](https://go.dev/doc/install).

## Layout

- One `go.mod` at the root; each example is its own `main` package in its own
  directory under `examples/`.
- Source files are named after their example (`c01-hello/hello.go`), not `main.go`.

## Resources & credit

Structure inspired by
[Jeremy Chone's xp repos](https://github.com/jeremychone-channel).