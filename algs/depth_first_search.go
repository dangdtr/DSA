package algs

type DepthFirstSearch struct {
	marked []bool //marked[v] = is there an s-v path?
	count  int    //count = number of vertices connected to s
}

// NewDepthFirstSearch G the graph
// s the source vertex
func NewDepthFirstSearch(G *Graph, s int) *DepthFirstSearch {
	marked := make([]bool, G.V())
	dfs := &DepthFirstSearch{marked: marked}
	dfs.DFS(G, s)
	return dfs
}

// DFS ...
func (s *DepthFirstSearch) DFS(G *Graph, v int) {
	s.count++
	s.marked[v] = true
	for _, w := range G.AdjInt(v) {
		if !s.marked[w] {
			s.DFS(G, w)
		}
	}
}

// Marked ...
func (s *DepthFirstSearch) Marked(v int) bool {
	return s.marked[v]
}

// Count ...
func (s *DepthFirstSearch) Count() int {
	return s.count
}
