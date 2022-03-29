package algs

import (
	"math/rand"
	"time"
)

// IntSliceToInterface Convert
//  []int to []interface{}
func IntSliceToInterface(slice []int) []interface{} {
	unsortedSlice := make([]interface{}, 0)
	for _, x := range slice {
		unsortedSlice = append(unsortedSlice, x)
	}
	return unsortedSlice
}

// StringSliceToInterface Convert
//  []string to []interface{}
func StringSliceToInterface(slice []string) []interface{} {
	unsortedSlice := make([]interface{}, 0)
	for _, x := range slice {
		unsortedSlice = append(unsortedSlice, x)
	}
	return unsortedSlice
}

// GenerateInterfaceSlice Generate Int
func GenerateInterfaceSlice(size int) []interface{} {

	slice := make([]interface{}, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// slice[i] = rand.Intn(999) - rand.Intn(999)
		slice[i] = rand.Intn(100) - rand.Intn(50)

	}
	return slice
}

// GenerateStringSlice Generate string
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
