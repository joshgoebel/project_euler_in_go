// Problem 5 - http://projecteuler.net/
//
// What is the smallest number that is evenly divisible by
// all of the numbers from 1 to 20?

package main

import (
	"fmt";
)

func divisible(n int64) bool {
	// we can do < 20 because we increment in units of 20 in the
	// main loop so we don't need to check divisibility by 20
	for i := int64(1); i < 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true;
}

func main() {
	best := int64(20);
	for {
		if divisible(best) {
			break
		}
		best += 20;
	}
	fmt.Printf("%d\n", best);
}
