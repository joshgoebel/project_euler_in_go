// Problem 2 - http://projecteuler.net/
//
// Find the sum of all the even-valued terms in Fib
// the sequence which do not exceed four million.

package main

import "fmt"

func main() {
	sum := 0;
	for a, b := 1, 2; b < 4000000; {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, b+a;
	}
	fmt.Printf("%d\n", sum);
}
