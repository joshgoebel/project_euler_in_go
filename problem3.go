// Problem 3 - http://projecteuler.net/
//
// Find the largest prime factor of a composite number

package main

import "fmt"

func prime(number int64) bool {
	for i := int64(2); i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true;
}

func main() {
	largest := int64(0);
	opposite := int64(0);
	composite := int64(600851475143);
	for i := int64(2); i < composite; i++ {
		if composite%i == 0 {
			opposite = composite / i;
			if prime(opposite) {
				largest = opposite;
				break;
			}
		}
	}
	fmt.Printf("%d\n", largest);
}
