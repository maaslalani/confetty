package array

import "math/rand"

// Sample returns a random element from a generic array
func Sample[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}
