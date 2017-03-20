package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	f := fire()
	c := factorial(f)
	for n := range c {
		fmt.Println(n)
	}
}

func fire() chan uint64 {
	out := make(chan uint64)
	go func() {
		var uu uint64 = 10000
		for i := 0; i < 50; i++ {
			for j := uu; j > 0; j-- {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n uint64) uint64 {
	var out uint64 = 1
	for i := n; n > 0; n-- {
		out *= i
	}
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/
