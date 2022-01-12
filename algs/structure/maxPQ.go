package structure

// import "fmt"

// type MaxPQ struct {
// 	n  int
// 	pq []PQItem
// }
type MaxPQ struct {
		n  int
		pq []interface{}
	}

type PQItem struct {
	value interface{}
}

// type PQItem interface {
// 	fmt.Stringer
// 	CompareTo(interface{}) int
// }

func NewMaxPQ(cap int) *MaxPQ {
	pq := make([]interface{}, cap+1)
	return &MaxPQ{pq: pq}
}



func NewMaxPQFrom(keyPQ []interface{}) *MaxPQ {
	pqKey := make([]interface{}, len(keyPQ)+1)
	
	pq := &MaxPQ{pq: pqKey}

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

// func (pq *MaxPQ) less(i, j int) bool {
// 	cmp := pq.pq[i].CompareTo(pq.pq[j])
// 	return cmp < 0
// }

func (pq *MaxPQ) less(i, j int) bool {
	a := pq.pq[i]
	b := pq.pq[j]
	switch a.(type) {
	case int:
		if ai, ok := a.(int); ok {
			if bi, ok := b.(int); ok {
				return ai < bi
			}
		}
	case string:
		if ai, ok := a.(string); ok {
			if bi, ok := b.(string); ok {
				return ai < bi
			}
		}
	default:
		panic("Unknown")
	}
	return false
}

func (pq *MaxPQ) exch(i, j int) {
	pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i]
}

func (pq *MaxPQ) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *MaxPQ) sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *MaxPQ) isEmpty() bool {
	return pq.n == 0
}

func (pq *MaxPQ) Insert(key PQItem) {
	if pq.n == len(pq.pq)-1 {
		pq.resize(2 * len(pq.pq))
	}

	pq.n++
	pq.pq[pq.n] = key
	pq.swim(pq.n)

	//isMaxHeap();
}

func (pq *MaxPQ) DelMax() interface{} {
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

func (pq *MaxPQ) Max() interface{} {
	return pq.pq[1]
}

func (pq *MaxPQ) resize(cap int) {
	temp := make([]interface{}, cap)
	for i := 1; i <= pq.n; i++ {
		temp[i] = pq.pq[i]
	}
	pq.pq = temp
}

func (pq *MaxPQ) Show() (in []interface{}) {
	for i := 0; i <= pq.n;i++ {
		in = append(in, pq.pq[i])
	}
	return
}