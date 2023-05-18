package dag

import "fmt"

// Rank returns the vertices in the r-th rank.
// Rank(0) equals "vertices without incoming edges".
// d.Rank(r) equals to {d - vertices without incoming edges}.Rank(r-1)
//
// For the following DAG:
//
//	a -> b -> c
//	 \   d -/
//	  > e
//
// Rank(0) returns [a, d] // [a, d] has no incoming edges
// Rank(1) returns [b, e] // remove [a, d], then [b, e] has no incoming edges
// Rank(2) returns [c]
func (d *DAG) Rank(r int) ([]Vertex, error) {
	vertexLen := len(d.vertices.ToSlice())
	if r < 0 || r >= vertexLen { // rank is [0, vertexLen-1], right-end means a single chain
		return nil, nil
	}
	if len(d.rank) < vertexLen {
		d.rank = make([][]Vertex, vertexLen)
	}
	if d.IsSorted() {
		return d.rank[r], nil
	}
	d.rank[0] = d.getVerticesNoIn()
	if len(d.rank[0]) == 0 && vertexLen > 0 { // still have vertices, but all have incoming edges
		return nil, fmt.Errorf("cycle detected")
	}
	if r == 0 { // exit condition
		return d.rank[0], nil
	}
	// dn is the graph removing vertices in d.rank[0]
	dn := d.Clone()
	dn.rank = d.rank[1:]
	dn.DeleteVertex(d.rank[0]...)
	if _, err := dn.Rank(r - 1); err != nil {
		return nil, err
	}
	return d.rank[r], nil
}

// getVerticesNoIn returns vertices without incoming edges.
func (d *DAG) getVerticesNoIn() []Vertex {
	v := d.vertices.Clone()
	for _, e := range d.edges.ToSlice() {
		v.RemoveAll(e.To())
	}
	return v.ToSlice()
}

// IsSorted checks whether the DAG is sorted (all ranks are computed).
func (d *DAG) IsSorted() bool {
	rankedVertex := 0
	for _, r := range d.rank {
		rankedVertex += len(r)
	}
	return len(d.vertices.ToSlice()) == rankedVertex
}

// Sort returns the vertices in batch, each batch contains vertices in the same rank.
func (d *DAG) Sort() ([][]Vertex, error) {
	_, err := d.Rank(len(d.vertices.ToSlice()) - 1)
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
