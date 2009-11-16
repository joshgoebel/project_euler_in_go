// Problem 16 - http://projecteuler.net/
//
// What is the first term in the Fibonacci sequence to contain 1000 digits?

package main

import (
	"fmt";
	"bignum";
)

func main() {
	a := bignum.Int(1);
	b := bignum.Int(1);
	i := 2;
	for {
		a, b = b, a.Add(b);
		i += 1;
		if len(b.String()) == 1000 {
			break
		}

	}
	fmt.Printf("%d\n", i);
}
