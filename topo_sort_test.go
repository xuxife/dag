package dag_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xuxife/dag"
)

func TestDAG(t *testing.T) {
	newVertices := func() (dag.Vertex, dag.Vertex, dag.Vertex, dag.Vertex, dag.Vertex) {
		return dag.NewVertex("a"), dag.NewVertex("b"), dag.NewVertex("c"), dag.NewVertex("d"), dag.NewVertex("e")
	}
	a, b, c, d, e := newVertices()
	//     /-- a
	//    /--- b
	// --|---- c
	//    \--- d
	//     \-- e
	fanOut, err := dag.New(
		dag.Add(a, b, c, d, e),
		dag.Add(a, b, c, d, e), // dag.New is idempotent
	)
	require.NoError(t, err)
	// a -> b -> c -> d -> e
	chain, err := dag.New(
		dag.From(a).To(b),
		dag.From(b).To(c),
		dag.From(c).To(d),
		dag.From(d).To(e),
	)
	require.NoError(t, err)
	// a ---> b ---> c
	// ^             |
	//  \           /
	//   e <-- d <-
	cycle, err := dag.New(
		dag.From(a).To(b),
		dag.From(b).To(c),
		dag.From(c).To(d),
		dag.From(d).To(e),
		dag.From(e).To(a),
	)
	require.ErrorContains(t, err, "cycle detected")
	_ = cycle
	// a ---> b --> d ----> e
	//   \         /
	//    --> c <-
	flag, err := dag.New(
		dag.From(a).To(b),
		dag.From(b).To(d),
		dag.From(a).To(c),
		dag.From(d).To(c),
		dag.From(d).To(e),
	)
	require.NoError(t, err)
	//	a ---> b ---> c
	//  |      ^	  ^
	//	 \    / \     |
	//	  -> e   d --/
	glass, err := dag.New(
		dag.From(a).To(b),
		dag.From(a).To(e),
		dag.From(b).To(c),
		dag.From(d).To(c),
		dag.From(e).To(b),
		dag.From(d).To(b),
	)
	require.NoError(t, err)
	for name, tc := range map[string]struct {
		d      *dag.DAG
		Expect [][]dag.Vertex
	}{
		"fanOut": {fanOut, [][]dag.Vertex{
			{a, b, c, d, e},
		}},
		"chain": {chain, [][]dag.Vertex{
			{a}, {b}, {c}, {d}, {e},
		}},
		"flag": {flag, [][]dag.Vertex{
			{a}, {b}, {d}, {c, e},
		}},
		"glass": {glass, [][]dag.Vertex{
			{a, d}, {e}, {b}, {c},
		}},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for i := range tc.Expect {
				actual, err := tc.d.Rank(i)
				require.NoError(t, err)
				require.ElementsMatch(t, tc.Expect[i], actual)
			}
		})
	}
}
