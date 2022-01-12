package sort


func Less(a, b interface{}) bool {
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


func IsSorted(a []interface{}) bool {
	return IsSortedLoHi(a, 0, len(a)-1)
}

func IsSortedLoHi(a []interface{}, lo int, hi int) bool {
	for i := lo + 1; i <= hi; i++ {
		if Less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}

