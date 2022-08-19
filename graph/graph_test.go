package graph

import (
	"testing"
)

func Test_graph_Topological(t *testing.T) {
	g := NewDirected()
	g.AddEdge(3, 4)
	g.AddEdge(2, 3)
	g.AddEdge(1, 2)
	order := g.Topological()

	t.Log(order)
	if len(order) != 4 {
		t.Errorf("order's length must be %v", 4)
	}

	for idx, vertex := range order {
		if vertex != idx+1 {
			t.Errorf("order[%v](%v) must be %v", idx, vertex, idx+1)
		}
	}
}
