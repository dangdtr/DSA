package algs

// NonRecursiveDFS DFS using stack.
type NonRecursiveDFS struct {
	marked []bool //marked[v] = is there an s-v path?
	trace  []int  //trace of dfs
}

// NewNonRecursiveDFS ...
func NewNonRecursiveDFS(G *Graph, s int) *NonRecursiveDFS {
	dfs := &NonRecursiveDFS{marked: make([]bool, G.V())}
	stack := NewStack[int]()
	stack.Push(s)
	dfs.marked[s] = true
	//fmt.Printf("%d |", s) // trace of dfs
	dfs.trace = append(dfs.trace, s)

	for !stack.IsEmpty() {
		//fmt.Print("Stack: ")
		// stack trace (optional)
		u := stack.Pop()
		if !dfs.Marked(u) {
			//fmt.Printf("%d |", u) // trace of dfs
			dfs.trace = append(dfs.trace, u)

			dfs.marked[u] = true
		}
		reverse := ReverseInt(G.AdjInt(u)) // reverse adj to pop smaller vertex first
		for _, v := range reverse {
			if !dfs.Marked(v) {
				stack.Push(v)
			}
		}
	}
	return dfs
}

// Trace ...
func (s *NonRecursiveDFS) Trace() []int {
	return s.trace
}

// Marked ...
func (s *NonRecursiveDFS) Marked(v int) bool {
	return s.marked[v]
}

// HasUnmarkedVertexOf ...
func (s *NonRecursiveDFS) HasUnmarkedVertexOf(G *Graph, u int) bool {
	for _, w := range G.AdjInt(u) {
		if s.Marked(w) {
			return false
		}
	}
	return true
}
