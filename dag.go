package dag

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type DAG struct {
	vertices mapset.Set[Vertex]
	edges    mapset.Set[Edge]
	rank     [][]Vertex
}

func New(edges ...Edges) (*DAG, error) {
	d := &DAG{
		vertices: mapset.NewSet[Vertex](),
		edges:    mapset.NewSet[Edge](),
	}
	for _, es := range edges {
		for _, e := range es {
			f, t := e.From(), e.To()
			if f != nil {
				d.vertices.Add(f)
			}
			if t != nil {
				d.vertices.Add(t)
			}
			if f != nil && t != nil {
				d.edges.Add(e)
			}
		}
	}
	if _, err := d.Sort(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DAG) Clone() *DAG {
	return &DAG{
		vertices: mapset.NewSet(d.vertices.ToSlice()...),
		edges:    mapset.NewSet(d.edges.ToSlice()...),
		rank:     d.rank, // carry the reference
	}
}

func (d *DAG) Merge(ds ...*DAG) *DAG {
	for _, dd := range ds {
		d.vertices.Append(dd.vertices.ToSlice()...)
		d.edges.Append(dd.edges.ToSlice()...)
	}
	return d
}

func (d *DAG) AddVertex(vs ...Vertex) {
	d.vertices.Append(vs...)
}

func (d *DAG) AddEdge(es ...Edge) {
	d.edges.Append(es...)
	for _, e := range es {
		d.vertices.Append(e.From(), e.To())
	}
}

func (d *DAG) GetVertices() []Vertex {
	return d.vertices.ToSlice()
}

func (d *DAG) GetEdges() []Edge {
	return d.edges.ToSlice()
}

// GetEdgesFrom returns all edges outcoming from the vertex
func (d *DAG) GetEdgesFrom(v Vertex) []Edge {
	edges := []Edge{}
	for _, e := range d.edges.ToSlice() {
		if e.From() == v {
			edges = append(edges, e)
		}
	}
	return edges
}

// GetEdgesTo returns all edges incoming to the vertex
func (d *DAG) GetEdgesTo(v Vertex) []Edge {
	edges := []Edge{}
	for _, e := range d.edges.ToSlice() {
		if e.To() == v {
			edges = append(edges, e)
		}
	}
	return edges
}

// GetVerticesFrom returns all vertices having edges from the vertex
func (d *DAG) GetVerticesFrom(v Vertex) []Vertex {
	vs := []Vertex{}
	for _, e := range d.GetEdgesFrom(v) {
		vs = append(vs, e.To())
	}
	return vs
}

// GetVerticesTo returns all vertices having edges to the vertex
func (d *DAG) GetVerticesTo(v Vertex) []Vertex {
	vs := []Vertex{}
	for _, e := range d.GetEdgesTo(v) {
		vs = append(vs, e.From())
	}
	return vs
}

func (d *DAG) DeleteEdge(es ...Edge) {
	d.edges.RemoveAll(es...)
}

func (d *DAG) DeleteVertex(vs ...Vertex) {
	d.vertices.RemoveAll(vs...)
	for _, e := range d.edges.ToSlice() {
		for _, v := range vs {
			if e.From() == v || e.To() == v {
				d.edges.RemoveAll(e)
			}
		}
	}
}
