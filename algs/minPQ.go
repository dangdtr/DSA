package algs

// import "fmt"

// MinQP is a generic binary heap (min-heap) with a comparator.
type MinQP[T any] struct {
	n    int
	pq   []T // 1-indexed
	less func(a, b T) bool
}

// type PQItem interface {
// 	fmt.Stringer
// 	CompareTo(interface{}) int
// }

func NewMinQP[T any](capacity int, less func(a, b T) bool) *MinQP[T] {
	pq := make([]T, capacity+1)
	return &MinQP[T]{pq: pq, less: less}
}

func NewMinQPFrom[T any](keyPQ []T, less func(a, b T) bool) *MinQP[T] {
	pqKey := make([]T, len(keyPQ)+1)
	pq := &MinQP[T]{pq: pqKey, less: less}
	pq.n = len(keyPQ)
	for i := 0; i < pq.n; i++ {
		pq.pq[i+1] = keyPQ[i]
	}
	for k := pq.n / 2; k >= 1; k-- {
		pq.sink(k)
	}
	return pq
}

// func (pq *MinQP) greater(i, j int) bool {
// 	cmp := pq.pq[i].CompareTo(pq.pq[j])
// 	return cmp < 0
// }

func (pq *MinQP[T]) greater(i, j int) bool {
	// greater means pq.pq[i] > pq.pq[j]
	return pq.less(pq.pq[j], pq.pq[i])
}

func (pq *MinQP[T]) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

func (pq *MinQP[T]) swim(k int) {
	for k > 1 && pq.greater(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MinQP[T]) sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.greater(j, j+1) {
			j++
		}
		if !pq.greater(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *MinQP[T]) isEmpty() bool {
	return pq.n == 0
}

func (pq *MinQP[T]) Insert(key T) {
	if pq.n == len(pq.pq)-1 {
		pq.resize(2 * len(pq.pq))
	}
	pq.n++
	pq.pq[pq.n] = key
	pq.swim(pq.n)
}

func (pq *MinQP[T]) DelMin() T {
    minVal := pq.pq[1]
	pq.exch(1, pq.n)
	pq.n--
	pq.sink(1)
	var zero T
	pq.pq[pq.n+1] = zero
	if pq.n > 0 && pq.n == (len(pq.pq)-1)/4 {
		pq.resize(len(pq.pq) / 2)
	}
	//isMaxHeap()

    return minVal
}

func (pq *MinQP[T]) Min() T {
	return pq.pq[1]
}

func (pq *MinQP[T]) resize(newCap int) {
	temp := make([]T, newCap)
	for i := 1; i <= pq.n; i++ {
		temp[i] = pq.pq[i]
	}
	pq.pq = temp
}

func (pq *MinQP[T]) Show() (in []T) {
	for i := 0; i <= pq.n; i++ {
		in = append(in, pq.pq[i])
	}
	return
}
