// Problem 7 - http://projecteuler.net/
//
// What is the 10001st prime number?

package main

import (
  "fmt";
)

var prime_storage [1000]int64;
var primes []int64;

func inits()
{
  primes = prime_storage[0:1];
  primes[0]=2;
}

func prime(number int64) bool
{
  if number==2 { return true}
  l := int64(len(primes));
  i := int64(0);
  for ; i < l && primes[i]<number; i++ {
    if number%primes[i]==0 {
      return false;
    }
  }
  i = primes[i-1]+1;
  if i%2==1 { i-- }
  for ; i< number; i+=2 {
    if number%i==0 {
      return false;
    }
  }
  l=int64(len(primes));
  if number>primes[l-1] && l<1000 {
    primes = primes[0:l+1];
    primes[l]=number
  }
  return true;
}

func main()
{
  inits();
  num := int64(1);
  for i:=0; i<10001; {
    num++;
    if prime(num) {
      i++
    }
  }
  fmt.Printf("%d\n", num);
//  for in, v := range primes {
//    fmt.Printf("prime [%d]: %d\n", in, v);
//  }
}