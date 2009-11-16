// Problem 12 - http://projecteuler.net/
//

// What is the value of the first triangle number 
// to have over five hundred divisors?

package main

import (
	"fmt";
	"./primes";
	// "math";
)


func main()
{
	triangle := int64(1);
	c := int64(0);
	for i:=triangle+1;; i++ {
		triangle += i;
		c = primes.Divisor_count(triangle);
		fmt.Printf("%d %d\n", triangle, c);
		if c > 500 {
			break
		}
	}
	fmt.Printf("%d\n", triangle);
}