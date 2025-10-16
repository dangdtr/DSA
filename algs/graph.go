package algs

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/dawnpanpan/go-dsa/stdin"
)

type Graph struct {
	v, e int
	adj  []*Bag[int]
}

func NewGraphWithV(V int) *Graph {
	adj := make([]*Bag[int], V)
	for i := 0; i < V; i++ {
		adj[i] = NewBag[int]()
	}
	return &Graph{v: V, e: 0, adj: adj}
}

func NewGraphWithG(G Graph) *Graph {
	adj := make([]*Bag[int], G.v)

	for i := 0; i < G.v; i++ {
		adj[i] = NewBag[int]()
	}

	for v := 0; v < G.v; v++ {
		items := G.adj[v].Iterator() // []int in current graph
		for i := len(items) - 1; i >= 0; i-- {
			adj[v].AddFirst(items[i])
		}
	}

	return &Graph{v: G.v, e: G.e, adj: adj}
}

func NewGraphWithFile(in *stdin.In) *Graph {
	v := in.ReadInt()
	adj := make([]*Bag[int], v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag[int]()
	}
	g := &Graph{v: v, e: 0, adj: adj}
	e := in.ReadInt()
	for i := 0; i < e; i++ {
		v, w := in.ReadInt(), in.ReadInt()
		g.AddEdge(v, w)
	}
	return g
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddEdge(v, w int) {
	g.e++
	g.adj[v].AddFirst(w)
	g.adj[w].AddFirst(v)
}

func (g *Graph) AdjInt(v int) (orderSlice []int) {
	temp := append([]int(nil), g.adj[v].Iterator()...)
	slices.Sort(temp)
	return temp
}

func (g *Graph) String() string {
	s := fmt.Sprintf("=================\n%d vertices, %d edges\n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		var adjs string
		for _, w := range g.AdjInt(i) {
			adjs = adjs + fmt.Sprintf(" %d", w)
		}
		s += fmt.Sprintf("%d: %s\n", i, adjs)
	}
	return s
}

func GraphGeneratorSimple(v, e int) *Graph {

	if e > (int)(v*(v-1)/2) {
		panic("too many edges")
	}
	if e < 0 {
		panic("too few edges")
	}
	G := NewGraphWithV(v)

	for G.e < e {
		V := rand.Intn(v)
		W := rand.Intn(v)
		G.AddEdge(V, W)

	}
	return G
}
