package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	// or letter := rune(word[0])
	// then -> func hashBucket(word string, buckets int32) int32 {
	// works

	bucket := letter % buckets
	return bucket
}
