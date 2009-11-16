// Problem 9 - http://projecteuler.net/
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import (
	"fmt";
	"math";
)

func main()
{
	var a, b, c int;
	var solved bool;
	for a = 1; a<1000; a++ {
		for b = a+1; b<1000; b++ {
			for c = b+1; c<1000; c++ {
				if a+b+c==1000 && ((math.Pow(float64(a),2)+math.Pow(float64(b),2)) == math.Pow(float64(c),2))
				{
					solved=true;
					break;
				}
			}
			if solved { break }
		}
		if solved { break }
	}
	
	fmt.Printf("total: %d\n", a+b+c);
	fmt.Printf("a: %d\n", a);
	fmt.Printf("b: %d\n", b);
	fmt.Printf("c: %d\n", c);
	fmt.Printf("product: %d\n", a*b*c);
}