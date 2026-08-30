# go-learn-basics

Small, self-contained Go programs one concept each, numbered in the order
learned, written from recall after study.

## Examples

| # | Example | Demonstrates |
|--:|---------|--------------|
| 01 | [hello](examples/c01-hello) | Toolchain proof: package main, imports, Println |
| 02 | [variables](examples/c02-variables) | var and :=, zero values, type conversion, untyped constants |
| 03 | [containers](examples/c03-containers) | Slices, maps, missing keys, range, slice aliasing |
| 04 | [functions](examples/c04-functions) | Functions, multiple returns, blank identifier, if/for/switch |
| 05 | [pointers](examples/c05-pointers) | Pointers, address-of, dereference, nil zero value |

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