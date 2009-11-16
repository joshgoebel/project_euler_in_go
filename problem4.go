// Problem 4 - http://projecteuler.net/
//
// Find the largest palindrome made from the product of 
// two 3-digit numbers.

package main

import (
  "fmt";
  "strconv";
  "strings";
)

func palidrome(s string) bool
{
  return reverse(s)==s;
}

func reverse(s string) string
{
  orig := strings.Bytes(s);
  len := len(s);
  n := make([]byte, len);
  for i:=0; i<len; i++ {
    n[len-i-1]=orig[i];
  }
  return string(n);
}

func main()
{
  best := 0;
  for i:=999; i>99; i-- {
    for j:=999; j>99; j-- {
      if palidrome(strconv.Itoa(j*i)) {
        if j*i > best {
          best=j*i;
        }
      }
    }
  }
  fmt.Printf("%d\n", best)
}