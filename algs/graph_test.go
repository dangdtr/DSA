package algs

import "testing"

func TestGraph_AddEdge_Adj_V_E_Copy_String(t *testing.T) {
	g := NewGraphWithV(5)
	if g.V() != 5 || g.E() != 0 {
		t.Fatalf("initial graph V=%d E=%d", g.V(), g.E())
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(3, 4)
	if g.E() != 3 {
		t.Fatalf("E()=%d, want 3", g.E())
	}

	adj0 := g.AdjInt(0)
	// AdjInt returns sorted neighbors
	if len(adj0) != 2 || adj0[0] != 1 || adj0[1] != 2 {
		t.Fatalf("AdjInt(0)=%v, want [1 2]", adj0)
	}

	// Copy
	g2 := NewGraphWithG(*g)
	g.AddEdge(1, 2)
	if g2.E() != 3 || g.E() != 4 {
		t.Fatalf("copy should be independent: g2.E=%d g.E=%d", g2.E(), g.E())
	}

	// Smoke test String()
	_ = g.String()
}


