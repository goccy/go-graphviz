package graphviz

type GraphOption func(g *Graphviz)

func WithName(name string) GraphOption {
	return func(g *Graphviz) {
		g.name = name
	}
}

func WithDirectedType(desc *GraphDescriptor) GraphOption {
	return func(g *Graphviz) {
		g.dir = desc
	}
}
