package dag

type Edges []Edge

type Edge interface {
	From() Vertex
	To() Vertex
}

type BaseEdge struct {
	from Vertex
	to   Vertex
}

func (e *BaseEdge) From() Vertex {
	return e.from
}

func (e *BaseEdge) To() Vertex {
	return e.to
}

var Add = From

func From(vs ...Vertex) Edges {
	rv := Edges{}
	for _, v := range vs {
		rv = append(rv, &BaseEdge{v, nil})
	}
	return rv
}

func (es Edges) To(to ...Vertex) Edges {
	edges := make([]Edge, 0, len(es)*len(to))
	for _, e := range es {
		f := e.From()
		for _, t := range to {
			edges = append(edges, &BaseEdge{f, t})
		}
	}
	return edges
}
