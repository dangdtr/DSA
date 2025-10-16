package algs

import (
	"math/rand"
)

// GenerateIntSlice generates random ints.
func GenerateIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) - rand.Intn(50)
	}
	return slice
}

// GenerateStringSlice Generate string
func GenerateStringSlice(size int) []string {
	slice := make([]string, size, size)
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
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// ReverseInt reverse int slice
func ReverseInt(s []int) []int {
	reverse := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		reverse[(len(s)-1)-i] = s[i]
	}
	return reverse
}
