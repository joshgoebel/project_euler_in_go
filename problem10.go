// Problem 10 - http://projecteuler.net/
//
// Find the sum of all the primes below two million.

package main

import (
	"fmt";
	"./primes";
)


func main()
{
	// include the number 2 so we can start couting odd #s
	// sum := int64(2);
	// for i:=int64(3); i<2000000; i += 2 {
	// 	if primes.Prime(i)
	// 	{
	// 		// fmt.Printf("prime: %d\n", i);
	// 		sum += i;
	// 	}
	// }
	sum := int64(0) ;
	for _, v := range primes.Primes_up_to(2000000) {
		sum += v;
	}
	
	fmt.Printf("%d\n", sum);
}