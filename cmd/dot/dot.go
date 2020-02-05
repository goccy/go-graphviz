package main

import (
	"errors"
	"fmt"

	"github.com/goccy/go-graphviz"
	"github.com/jessevdk/go-flags"
)

type Option struct {
	Format     graphviz.Format `description:"specify output format ( currently supported: dot svg png jpg )" short:"T" default:"dot"`
	Layout     graphviz.Layout `description:"specify layout engine ( currently supported: circo dot fdp neato nop nop1 nop2 osage patchwork sfdp twopi )" short:"K" default:"dot"`
	OutputFile string          `description:"specify output file name" short:"o" required:"true"`
}

func _main(args []string, opt *Option) error {
	if len(args) == 0 {
		return errors.New("required dot file")
	}
	dotFile := args[0]
	graph, err := graphviz.ParseFile(dotFile)
	if err != nil {
		return err
	}
	g := graphviz.New()
	defer func() {
		graph.Close()
		g.Close()
	}()
	g.SetLayout(opt.Layout)
	return g.RenderFilename(graph, opt.Format, opt.OutputFile)
}

func main() {
	var opt Option
	parser := flags.NewParser(&opt, flags.Default)
	args, err := parser.Parse()
	if err != nil {
		return
	}
	if err := _main(args, &opt); err != nil {
		fmt.Println(err)
	}
}
