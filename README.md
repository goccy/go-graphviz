# go-graphviz [![GoDoc](https://godoc.org/github.com/goccy/go-graphviz?status.svg)](https://godoc.org/github.com/goccy/go-graphviz)

Go bindings for Graphviz ( port of version `2.40.1` )

# Features

- No need to install Graphviz library ( ~`brew install graphviz`~ or ~`apt-get install graphviz`~ )
- Supports parsing for DOT language
- Supports rendering graph in pure Go
- Supports switch renderer to your own
- Supports type safed property setting
- `gvc` `cgraph` are available as sub package

## Currently supported Layout

`circo` `dot` `fdp` `neato` `nop` `nop1` `nop2` `osage` `patchwork` `sfdp` `twopi`

## Currently supported format

`dot` `svg` `png` `jpeg`

# Installation

```bash
$ go get -u github.com/goccy/go-graphviz
```

# Synopsis

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
  graph := g.Graph()
  defer func() {
    graph.Close()
    g.Close()
  }()
  n := graph.CreateNode("n")
  m := graph.CreateNode("m")
  graph.CreateEdge("e", n, m).SetLabel("e")
  var buf bytes.Buffer
  if err := g.Render(graph, "dot", &buf); err != nil {
    log.Fatalf("%+v", err)
  }
  fmt.Println(buf.String())
}
```

# Tool

## `dot`

### Installation

```bash
$ go get -u github.com/goccy/go-graphviz/cmd/dot
```

### Usage

```
Usage:
  dot [OPTIONS]

Application Options:
  -T=         specify output format ( currently supported: dot svg png ) (default: dot)
  -o=         specify output file name

Help Options:
  -h, --help  Show this help message
```

# License

MIT
