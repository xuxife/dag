package dag

import "fmt"

type DAG struct {
	vertices Set[Vertex]
	edges    Set[Edge]
	rank     [][]Vertex
}

func New(edges ...Edges) (*DAG, error) {
	d := &DAG{
		vertices: NewSet[Vertex](),
		edges:    NewSet[Edge](),
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
	return d, nil
}

func (d *DAG) Duplicate() *DAG {
	return &DAG{
		vertices: NewSet(d.vertices.List()...),
		edges:    NewSet(d.edges.List()...),
		rank:     d.rank, // carry the reference
	}
}

func (d *DAG) Merge(ds ...*DAG) *DAG {
	for _, dd := range ds {
		d.vertices.Add(dd.vertices.List()...)
		d.edges.Add(dd.edges.List()...)
	}
	return d
}

func (d *DAG) AddVertex(vs ...Vertex) {
	d.vertices.Add(vs...)
}

func (d *DAG) AddEdge(es ...Edge) {
	d.edges.Add(es...)
	for _, e := range es {
		d.vertices.Add(e.From(), e.To())
	}
}

func (d *DAG) GetVertices() []Vertex {
	return d.vertices.List()
}

func (d *DAG) GetEdges() []Edge {
	return d.edges.List()
}

func (d *DAG) GetEdgesFrom(v Vertex) []Edge {
	edges := []Edge{}
	for _, e := range d.edges.List() {
		if e.From() == v {
			edges = append(edges, e)
		}
	}
	return edges
}

func (d *DAG) DeleteEdge(es ...Edge) {
	d.edges.Delete(es...)
}

func (d *DAG) DeleteVertex(vs ...Vertex) {
	d.vertices.Delete(vs...)
	for _, e := range d.edges.List() {
		for _, v := range vs {
			if e.From() == v || e.To() == v {
				d.edges.Delete(e)
			}
		}
	}
}

func (d *DAG) Rank(r int) ([]Vertex, error) {
	if len(d.rank) < d.vertices.Len() {
		d.rank = make([][]Vertex, d.vertices.Len())
	}
	if r < 0 || r >= d.vertices.Len() {
		return nil, nil
	}
	if d.IsSorted() {
		return d.rank[r], nil
	}
	d.rank[0] = d.getRank0Vertices()
	if len(d.rank[0]) == 0 && d.vertices.Len() > 0 {
		return nil, fmt.Errorf("cycle detected")
	}
	if r == 0 {
		return d.rank[0], nil
	}
	dn := d.Duplicate()
	dn.rank = d.rank[1:]
	dn.DeleteVertex(d.rank[0]...)
	var err error
	if d.rank[1], err = dn.Rank(r - 1); err != nil {
		return nil, err
	}
	return d.rank[r], nil
}

func (d *DAG) getRank0Vertices() []Vertex {
	v := d.vertices.Duplicate()
	for _, e := range d.edges.List() {
		v.Delete(e.To())
	}
	return v.List()
}

func (d *DAG) IsSorted() bool {
	rankedVLen := 0
	for _, r := range d.rank {
		rankedVLen += len(r)
	}
	return d.vertices.Len() == rankedVLen
}

func (d *DAG) Sort() ([][]Vertex, error) {
	_, err := d.Rank(d.vertices.Len() - 1)
	if err != nil {
		return nil, err
	}
	rv := [][]Vertex{}
	for _, batch := range d.rank {
		if len(batch) > 0 {
			rv = append(rv, batch)
		}
	}
	return rv, nil
}
