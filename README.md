[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-908a85?logo=gitpod)](https://gitpod.io/from-referrer/)

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

## Information

Basic Golang translations of the [Princeton - Java code implementation](http://algs4.cs.princeton.edu/code/)
in the textbook Algorithms, 4th Edition by Robert Sedgewick and Kevin Wayne.

## Reference

https://github.com/shellfly/algo

## Project overview

This project is a collection of classic data structures and algorithms implemented in Go. It aims to provide clear, idiomatic Go translations suitable for learning, practice, and use as building blocks in other programs.

### Features

- Sorting: insertion, mergesort (top-down and bottom-up), quicksort, 3-way quicksort
- Core DS: Bag, Stack, Queue, Singly Linked List, Priority Queues (Min/Max), Binary Search Tree (BST)
- Graph: undirected graph representation, BFS and DFS (recursive and non-recursive), shortest paths by edge count
- Simple file-based input readers for demos (e.g., graph files)

### Requirements

- Go 1.23+

### Directory structure

- `algs/`: algorithms and data structures
  - sorting: `insertionSort.go`, `mergesort.go`, `mergesort_bottom_up.go`, `quicksort.go`, `quicksort3way.go`
  - structures: `bag.go`, `stack.go`, `queue.go`, `linkedlist/`
  - priority queues: `maxPQ.go`, `minPQ.go`
  - trees: `bst.go`
  - graphs: `graph.go`, `breath_first_paths.go`, `depth_first_search.go`, `depth_first_search_lifo.go`, `depth_first_path.go`
- `cmd/`: small demo programs and sample inputs

## Getting started

```bash
go mod tidy
cd cmd && go run .
```

Run from `cmd/` if a demo expects files (e.g., `vnoi-bfs.txt`) so relative paths resolve.

## Usage examples

### Sorting (comparator style)

```go
import (
    "cmp"
    "github.com/dawnpanpan/go-dsa/algs"
)

nums := []int{5, 3, 8, 1}
algs.QuickSortFunc(nums, cmp.Less[int])
```

### Priority queue

```go
import (
    "cmp"
    "github.com/dawnpanpan/go-dsa/algs"
)

pq := algs.NewMaxPQ[int](128, cmp.Less[int])
pq.Insert(10)
pq.Insert(7)
_ = pq.DelMax() // 10
```

### BST

```go
import "github.com/dawnpanpan/go-dsa/algs"

bst := algs.NewBST[string, string]()
bst.Put("A", "alpha")
if v, ok := bst.Get("A"); ok {
    _ = v
}
```

### Graph input format

- Line 1: number of vertices `V`
- Line 2: number of edges `E`
- Next `E` lines: undirected edges `v w`

Example (`cmd/vnoi-bfs.txt`):

```
15
15
1 3
1 2
1 4
...
```

## Acknowledgements

- Princeton Algorithms, 4th Edition (Sedgewick & Wayne) and the [algs4 code](http://algs4.cs.princeton.edu/code/)
- Community resources such as [shellfly/algo](https://github.com/shellfly/algo)

This project is a respectful adaptation for Go, keeping the spirit and educational value of the originals.
