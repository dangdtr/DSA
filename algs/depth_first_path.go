package algs

type DepthFirstPaths struct {
	marked []bool // marked[v] = is there an s-v path?
	edgeTo []int  // edgeTo[v] = last edge on s-v path
	s      int    // source vertex
}

func NewDepthFirstPaths(G *Graph, s int) *DepthFirstPaths {
	dfs := &DepthFirstPaths{marked: make([]bool, G.V()), edgeTo: make([]int, G.V()), s: s}
	dfs.DFS(G, s)
	return dfs
}

func (p *DepthFirstPaths) DFS(G *Graph, v int) {
	p.marked[v] = true
	for _, w := range G.AdjInt(v) {
		if !p.marked[w] {
			p.edgeTo[w] = v
			p.DFS(G, w)
		}
	}
}

func (p *DepthFirstPaths) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *DepthFirstPaths) PathTo(v int) []int {
	if !p.HasPathTo(v) {
		return nil
	}

	path := NewStack()
	for x := v; x != p.s; x = p.edgeTo[x] {
		path.Push(x)
	}
	path.Push(p.s)
	return ReverseInt(path.IteratorIntSlide())
}
