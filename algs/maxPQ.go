package algs

// import "fmt"

// MaxPQ is a generic binary heap (max-heap) with a comparator.
type MaxPQ[T any] struct {
	n    int
	pq   []T // 1-indexed
	less func(a, b T) bool
}

func NewMaxPQ[T any](capacity int, less func(a, b T) bool) *MaxPQ[T] {
	pq := make([]T, capacity+1)
	return &MaxPQ[T]{pq: pq, less: less}
}

func NewMaxPQFrom[T any](keyPQ []T, less func(a, b T) bool) *MaxPQ[T] {
	pqKey := make([]T, len(keyPQ)+1)
	pq := &MaxPQ[T]{pq: pqKey, less: less}
	pq.n = len(keyPQ)
	for i := 0; i < pq.n; i++ {
		pq.pq[i+1] = keyPQ[i]
	}
	for k := pq.n / 2; k >= 1; k-- {
		pq.sink(k)
	}
	return pq
}

func (pq *MaxPQ[T]) lessIdx(i, j int) bool {
	return pq.less(pq.pq[i], pq.pq[j])
}

func (pq *MaxPQ[T]) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

func (pq *MaxPQ[T]) swim(k int) {
	for k > 1 && pq.lessIdx(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MaxPQ[T]) sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.lessIdx(j, j+1) {
			j++
		}
		if !pq.lessIdx(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *MaxPQ[T]) isEmpty() bool {
	return pq.n == 0
}

func (pq *MaxPQ[T]) Insert(key T) {
	if pq.n == len(pq.pq)-1 {
		pq.resize(2 * len(pq.pq))
	}
	pq.n++
	pq.pq[pq.n] = key
	pq.swim(pq.n)
}

func (pq *MaxPQ[T]) DelMax() T {
    maxVal := pq.pq[1]
	pq.exch(1, pq.n)
	pq.n--
	pq.sink(1)
	var zero T
	pq.pq[pq.n+1] = zero
	if pq.n > 0 && pq.n == (len(pq.pq)-1)/4 {
		pq.resize(len(pq.pq) / 2)
	}
    return maxVal
}

func (pq *MaxPQ[T]) Max() T {
	return pq.pq[1]
}

func (pq *MaxPQ[T]) resize(newCap int) {
	temp := make([]T, newCap)
	for i := 1; i <= pq.n; i++ {
		temp[i] = pq.pq[i]
	}
	pq.pq = temp
}

func (pq *MaxPQ[T]) Show() (in []T) {
	for i := 0; i <= pq.n; i++ {
		in = append(in, pq.pq[i])
	}
	return
}
