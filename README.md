# go-graphviz [![Go](https://github.com/goccy/go-graphviz/workflows/Go/badge.svg)](https://github.com/goccy/go-graphviz/actions) [![GoDoc](https://godoc.org/github.com/goccy/go-graphviz?status.svg)](https://pkg.go.dev/github.com/goccy/go-graphviz) 

Go bindings for Graphviz ( port of version `2.40.1` )

# Features

- No need to install Graphviz library ( ~`brew install graphviz`~ or ~`apt-get install graphviz`~ )
- Supports parsing for DOT language
- Supports rendering graph in pure Go
- Supports switch renderer to your own
- Supports type safed property setting
- `gvc` `cgraph` `cdt` are available as sub package

## Currently supported Layout

`circo` `dot` `fdp` `neato` `nop` `nop1` `nop2` `osage` `patchwork` `sfdp` `twopi`

## Currently supported format

`dot` `svg` `png` `jpg`

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
  "fmt"
  "log"

  "github.com/goccy/go-graphviz"
)

func main() {
  g := graphviz.New()
  graph, err := g.Graph()
  if err != nil {
    log.Fatal(err)
  }
  defer func() {
    if err := graph.Close(); err != nil {
      log.Fatal(err)
    }
    g.Close()
  }()
  n, err := graph.CreateNode("n")
  if err != nil {
    log.Fatal(err)
  }
  m, err := graph.CreateNode("m")
  if err != nil {
    log.Fatal(err)
  }
  e, err := graph.CreateEdge("e", n, m)
  if err != nil {
    log.Fatal(err)
  }
  e.SetLabel("e")
  var buf bytes.Buffer
  if err := g.Render(graph, "dot", &buf); err != nil {
    log.Fatal(err)
  }
  fmt.Println(buf.String())
}
```

## 2. Parse DOT Graph

```go
path := "/path/to/dot.gv"
b, err := ioutil.ReadFile(path)
if err != nil {
  log.Fatal(err)
}
graph := graphviz.ParseBytes(b)
```

## 3. Render Graph

```go
g := graphviz.New()
graph, err := g.Graph()
if err != nil {
  log.Fatal(err)
}

// create your graph

// 1. write encoded PNG data to buffer
var buf bytes.Buffer
if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
  log.Fatal(err)
}

// 2. get as image.Image instance
image, err := g.RenderImage(graph)
if err != nil {
  log.Fatal(err)
}

// 3. write to file directly
if err := g.RenderFilename(graph, graphviz.PNG, "/path/to/graph.png"); err != nil {
  log.Fatal(err)
}
```

# Tool

## `dot`

### Installation

```bash
$ go get github.com/goccy/go-graphviz/cmd/dot
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

<img width = "600px" src="https://user-images.githubusercontent.com/209884/75105919-48685b00-565c-11ea-8add-ebd5545f5399.png"></img>

`go-graphviz` has four layers.

1. `graphviz` package provides facade interface for manipulating all features of graphviz library
2. `gvc` `cgraph` `cdt` are sub packages ( FYI: C library section in https://www.graphviz.org/documentation )
3. `internal/ccall` package provides bridge interface between Go and C
4. `go-graphviz` includes full graphviz sources

# License

MIT
