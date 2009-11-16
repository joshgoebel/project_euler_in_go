// Problem 16 - http://projecteuler.net/
//
// What is the sum of the digits of the number 21000?

// Ruby
// puts (2**1000).to_s.split(//).map {|x| x.to_i}.inject(0) {|sum,n| sum+n }

package main

import (
	"fmt";
	"bignum";
	"strconv";
)

func main() {
	sum := 0;
	two_1000 := bignum.Int(2);
	for i := 1; i < 1000; i++ {
		two_1000 = two_1000.Mul(bignum.Int(2))
	}
	for _, v := range two_1000.String() {
		c, _ := strconv.Atoi(string(v));
		sum += c;
	}
	fmt.Printf("%d\n", sum);
}
