package sort

import (
	"math/rand"
	"time"
)

func GenerateInterfaceSlice(size int) []interface{} {

	slice := make([]interface{}, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice[i] = rand.Intn(999) - rand.Intn(999)
		slice[i] = rand.Intn(100) - rand.Intn(50)

	}
	return slice
}

func GenerateStringSlice(size int) []string {
	slice := make([]string, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// n := rand.Intn(52) - rand.Intn(0)
		n := rand.Intn(10)
		slice[i] = generateString(n)
	}
	return slice

}
func generateString(size int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, size)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func InitEmptyInterfaceSlice(size int) []interface{} {

	slice := make([]interface{}, size, size)
	return slice
}

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

func LessInt(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
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

// func GenerateSlice(size int) []int {

// 	slice := make([]int, size, size)
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < size; i++ {
// 		slice[i] = rand.Intn(999) - rand.Intn(999)
// 	}
// 	return slice
// }

// func InitEmptySlice(size int) []int {

// 	slice := make([]int, size, size)
// 	return slice
// }
