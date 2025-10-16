package algs

import (
	"math"
	"slices"
)

const INFINITY = math.MaxInt

type BreathFirstPaths struct {
	marked []bool // marked[v] = is there an s-v path?
	edgeTo []int  // edgeTo[v] = last edge on s-v path
	distTo []int  // distTo[v] = number of edges shortest s-v path}
	trace  []int  //trace = trace of bfs
}

func NewBreathFirstPaths(G *Graph, s int) *BreathFirstPaths {
	bfs := &BreathFirstPaths{marked: make([]bool, G.V()), edgeTo: make([]int, G.V()), distTo: make([]int, G.V())}
	bfs.BFS(G, s)
	return bfs
}

func (p *BreathFirstPaths) BFS(G *Graph, v int) {
	queue := NewQueue[int]()
	for v := 0; v < G.V(); v++ {
		p.distTo[v] = INFINITY
	}
	p.distTo[v] = 0
	p.marked[v] = true
	p.trace = append(p.trace, v)

	queue.Enqueue(v)
	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for _, w := range G.AdjInt(v) {
			if !p.marked[w] {
				p.distTo[w] = p.distTo[v] + 1
				p.marked[w] = true
				p.edgeTo[w] = v
				queue.Enqueue(w)
				p.trace = append(p.trace, w)
			}
		}
	}
}

func (p *BreathFirstPaths) Trace() []int {
	return p.trace
}

func (p *BreathFirstPaths) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *BreathFirstPaths) PathTo(v int) []int {
	if !p.HasPathTo(v) {
		return nil
	}

	var path []int
	for x := v; p.distTo[x] != 0; x = p.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, p.sourceOf(v))
	slices.Reverse(path)
	return path
}

// sourceOf finds the source vertex for a node on a BFS tree by rewinding distTo.
func (p *BreathFirstPaths) sourceOf(v int) int {
	x := v
	for p.distTo[x] != 0 {
		x = p.edgeTo[x]
	}
	return x
}
