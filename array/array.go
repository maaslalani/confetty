package array

import "math/rand"

// Sample returns a random element from an array
// (can't wait for generics!)
func Sample(s []string) string {
	return s[rand.Intn(len(s))]
}
