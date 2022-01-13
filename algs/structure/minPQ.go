package structure

// import "fmt"

// type MinQP struct {
// 	n  int
// 	pq []PQItem
// }
type MinQP struct {
		n  int
		pq []interface{}
	}


// type PQItem interface {
// 	fmt.Stringer
// 	CompareTo(interface{}) int
// }

func NewMinQP(cap int) *MinQP {
	pq := make([]interface{}, cap+1)
	return &MinQP{pq: pq}
}



func NewMinQPFrom(keyPQ []interface{}) *MinQP {
	pqKey := make([]interface{}, len(keyPQ)+1)
	
	pq := &MinQP{pq: pqKey}

	pq.n = len(keyPQ)
	
	for i := 0; i < pq.n; i++ {
		pq.pq[i+1] = keyPQ[i]

	}
	for k := pq.n/2; k >= 1; k-- {
		pq.sink(k)
	}
	//isMaxHeap();
	return pq
}

// func (pq *MinQP) greater(i, j int) bool {
// 	cmp := pq.pq[i].CompareTo(pq.pq[j])
// 	return cmp < 0
// }

func (pq *MinQP) greater(i, j int) bool {
	a := pq.pq[i]
	b := pq.pq[j]
	switch a.(type) {
	case int:
		if ai, ok := a.(int); ok {
			if bi, ok := b.(int); ok {
				return ai > bi
			}
		}
	case string:
		if ai, ok := a.(string); ok {
			if bi, ok := b.(string); ok {
				return ai > bi
			}
		}
	default:
		panic("Unknown")
	}
	return false
}

func (pq *MinQP) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

func (pq *MinQP) swim(k int) {
	for k > 1 && pq.greater(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MinQP) sink(k int) {
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

func (pq *MinQP) isEmpty() bool {
	return pq.n == 0
}

func (pq *MinQP) Insert(key PQItem) {
	if pq.n == len(pq.pq)-1 {
		pq.resize(2 * len(pq.pq))
	}

	pq.n++
	pq.pq[pq.n] = key
	pq.swim(pq.n)

	//isMaxHeap();
}

func (pq *MinQP) DelMin() interface{} {
	max := pq.pq[1]
	pq.exch(1, pq.n)
	pq.n--
	pq.sink(1)
	pq.pq[pq.n+1] = nil
	if pq.n > 0 && pq.n == (len(pq.pq)-1)/4 {
		pq.resize(len(pq.pq) / 2)
	}
	//isMaxHeap()

	return max
}

func (pq *MinQP) Min() interface{} {
	return pq.pq[1]
}

func (pq *MinQP) resize(cap int) {
	temp := make([]interface{}, cap)
	for i := 1; i <= pq.n; i++ {
		temp[i] = pq.pq[i]
	}
	pq.pq = temp
}

func (pq *MinQP) Show() (in []interface{}) {
	for i := 0; i <= pq.n;i++ {
		in = append(in, pq.pq[i])
	}
	return
}