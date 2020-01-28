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
  "github.com/goccy/go-graphviz"
)

func main() {
  g := graphviz.New()
  graph := g.Graph()
  defer func() {
    graph.Close()
    g.Close()
  }()
  n1 := graph.Node("n1")
  n2 := graph.Node("n2")
  n3 := graph.Node("n3")
  n1.SetShape("diamond")
  e1 := graph.Edge("e1", n1, n2)
  e1.SetLabel("e1")
  graph.Edge("e2", n1, n3)
  g.RenderFilename(graph, "png", "sample.png")
}
```

# Tool

## `dot`

### Installation

```bash
$ go get -u github.com/goccy/go-graphviz/cmd/dot
```

### Usage

# License

MIT
