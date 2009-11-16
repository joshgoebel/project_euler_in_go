// Problem 36 - http://projecteuler.net/
//
// Find the sum of all numbers, less than one million, 
// which are palindromic in base 10 and base 2.

package main

import (
	"fmt";
	"strconv";
	"strings";
)

func palidrome(s string) bool	{ return reverse(s) == s }

func reverse(s string) string {
	orig := strings.Bytes(s);
	len := len(s);
	n := make([]byte, len);
	for i := 0; i < len; i++ {
		n[len-i-1] = orig[i]
	}
	return string(n);
}

func main() {
	sum := 0;
	for i := 1; i < 1000000; i++ {
		if palidrome(strconv.Itoa(i)) &&
		   palidrome(strconv.Itob(i, 2)) {
			sum += i;
		}
	}
	fmt.Printf("%d\n", sum);
}
