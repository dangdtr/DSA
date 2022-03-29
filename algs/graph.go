package algs

import (
	"fmt"
	"github.com/dawnpanpan/go-dsa/stdin"
	"math/rand"
	"time"
)

type Graph struct {
	v, e int
	adj  []*Bag
}

func NewGraphWithV(V int) *Graph {
	adj := make([]*Bag, V)
	for i := 0; i < V; i++ {
		adj[i] = NewBag()
	}
	return &Graph{v: V, e: 0, adj: adj}
}

func NewGraphWithG(G Graph) *Graph {
	adj := make([]*Bag, G.v)

	for i := 0; i < G.v; i++ {
		adj[i] = NewBag()
	}

	for v := 0; v < G.v; v++ {
		reverse := NewStack()
		for _, w := range G.adj[v].InteratorSlide() {
			reverse.Push(w)
		}

		for _, w := range reverse.InteratorSlide() {
			adj[v].AddFirst(w)
		}
	}

	return &Graph{v: G.v, e: G.e, adj: adj}
}

func NewGraphWithFile(in *stdin.In) *Graph {
	v := in.ReadInt()
	adj := make([]*Bag, v)
	for i := 0; i < v; i++ {
		adj[i] = NewBag()
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

func (g *Graph) AdjInt(v int) []int {
	return g.adj[v].InteratorInt()
}

func (g *Graph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", g.v, g.e)
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
	rand.Seed(time.Now().UnixNano())

	G := NewGraphWithV(v)

	for G.e < e {
		V := rand.Intn(v)
		W := rand.Intn(v)
		G.AddEdge(V, W)

	}
	return G
}
