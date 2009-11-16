// Problem 10 - http://projecteuler.net/
//
// Find the sum of all the primes below two million.

package main

import (
	"fmt";
	"./primes";
)


func main() {
	sum := int64(0);
	for _, v := range primes.Primes_up_to(2000000) {
		sum += v
	}

	fmt.Printf("%d\n", sum);
}
