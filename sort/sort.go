package sort

import (
	"math/rand"
	"time"
)

// Convert
//  []int to []interface{}
func IntSliceToInterface(slice []int) []interface{} {
	unsortedSlice := make([]interface{}, 0)
	for _, x := range slice {
		unsortedSlice = append(unsortedSlice, x)
	}
	return unsortedSlice
}

// Convert
//  []string to []interface{}
func StringSliceToInterface(slice []string) []interface{} {
	unsortedSlice := make([]interface{}, 0)
	for _, x := range slice {
		unsortedSlice = append(unsortedSlice, x)
	}
	return unsortedSlice
}

// Generate Int 
func GenerateInterfaceSlice(size int) []interface{} {

	slice := make([]interface{}, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice[i] = rand.Intn(999) - rand.Intn(999)
		slice[i] = rand.Intn(100) - rand.Intn(50)

	}
	return slice
}

// Generate string 
func GenerateStringSlice(size int) []string {
	slice := make([]string, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// n := rand.Intn(52) - rand.Intn(0)
		n := rand.Intn(10)
		// n := 1
		slice[i] = generateString(n)
	}
	return slice

}
func generateString(size int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// var letterRunes = []rune("0123456789")

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

