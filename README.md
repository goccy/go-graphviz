# go-graphviz [![Go](https://github.com/goccy/go-graphviz/workflows/Go/badge.svg)](https://github.com/goccy/go-graphviz/actions) [![GoDoc](https://godoc.org/github.com/goccy/go-graphviz?status.svg)](https://pkg.go.dev/github.com/goccy/go-graphviz) 

Go bindings for Graphviz

<img src="https://user-images.githubusercontent.com/209884/90976476-64e84000-e578-11ea-9596-fb4a7d3b11a6.png" width="400px"></img>

# Features

Graphviz version is [here](./graphviz.version)

- Pure Go Library
- No need to install Graphviz library ( ~`brew install graphviz`~ or ~`apt-get install graphviz`~ )
  - The Graphviz library has been converted to WebAssembly (WASM) and embedded it, so it works consistently across all environments
- Supports encoding/decoding for DOT language
- Supports custom renderer for custom format
- Supports setting graph properties in a type-safe manner

## Supported Layout

`circo` `dot` `fdp` `neato` `nop` `nop1` `nop2` `osage` `patchwork` `sfdp` `twopi`

## Supported Format

`dot` `svg` `png` `jpg`

The above are the formats supported by default. You can also add custom formats.

# Installation

```bash
$ go get github.com/goccy/go-graphviz
```

# Synopsis

## 1. Write DOT Graph in Go

```go
package main

import (
  "bytes"
  "context"
  "fmt"
  "log"

  "github.com/goccy/go-graphviz"
)

func main() {
  ctx := context.Background()
  g, err := graphviz.New(ctx)
  if err != nil { panic(err )}

  graph, err := g.Graph()
  if err != nil { panic(err) }
  defer func() {
    if err := graph.Close(); err != nil { panic(err) }
    g.Close()
  }()
  n, err := graph.CreateNodeByName("n")
  if err != nil { panic(err) }

  m, err := graph.CreateNodeByName("m")
  if err != nil { panic(err) }

  e, err := graph.CreateEdgeByName("e", n, m)
  if err != nil { panic(err) }
  e.SetLabel("e")

  var buf bytes.Buffer
  if err := g.Render(ctx, graph, "dot", &buf); err != nil {
    log.Fatal(err)
  }
  fmt.Println(buf.String())
}
```

## 2. Parse DOT Graph

```go
path := "/path/to/dot.gv"
b, err := os.ReadFile(path)
if err != nil { panic(err) }
graph, err := graphviz.ParseBytes(b)
```

## 3. Render Graph

```go
ctx := context.Background()
g, err := graphviz.New(ctx)
if err != nil { panic(err) }

graph, err := g.Graph()
if err != nil { panic(err) }

// create your graph

// 1. write encoded PNG data to buffer
var buf bytes.Buffer
if err := g.Render(ctx, graph, graphviz.PNG, &buf); err != nil { panic(err) }

// 2. get as image.Image instance
image, err := g.RenderImage(ctx, graph)
if err != nil { panic(err) }

// 3. write to file directly
if err := g.RenderFilename(ctx, graph, graphviz.PNG, "/path/to/graph.png"); err != nil { panic(err) }
```

# Tool

## `dot`

### Installation

```bash
$ go install github.com/goccy/go-graphviz/cmd/dot@latest
```

### Usage

```
Usage:
  dot [OPTIONS]

Application Options:
  -T=         specify output format ( currently supported: dot svg png jpg ) (default: dot)
  -K=         specify layout engine ( currently supported: circo dot fdp neato nop nop1 nop2 osage patchwork sfdp twopi )
  -o=         specify output file name

Help Options:
  -h, --help  Show this help message
```

# How it works

1. Generates bindings between Go and C from [Protocol Buffers file](./internal/wasm/bind.proto).
2. Builds graphviz.wasm on the [docker container](./internal/wasm/build/Dockerfile).
3. Uses Graphviz functionality from a sub-packages ( `cdt` `cgraph` `gvc` ) via the `internal/wasm` package. 
4. `graphviz` package provides facade interface for all sub packages.

# License

MIT

This library embeds and uses `graphviz.wasm`, which is generated based on the original source code of Graphviz. Therefore, the `graphviz.wasm` follows [the license adopted by Graphviz](https://graphviz.org/license) ( Eclipse Public License ).
