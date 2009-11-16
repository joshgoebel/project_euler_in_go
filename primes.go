package primes

import (
	"once";
	"math";
	"fmt";
	"container/vector";
	)

var prime_storage [10000]int64;
var primes []int64;

func inits()
{
  primes = prime_storage[0:1];
  primes[0]=2;
}

func Prime(number int64) bool
{
  once.Do(inits);
  if number==2 { return true}
  l := int64(len(primes));
  i := int64(0);
  for ; i < l && primes[i]<number; i++ {
    if number%primes[i]==0 {
      return false;
    }
  }
  i = primes[i-1]+1;
  if i%2==0 { i-- }
  for ; i< number; i+=2 {
    if number%i==0 {
      return false;
    }
  }
  l=int64(len(primes));
  if number>primes[l-1] && l<10000 {
    primes = primes[0:l+1];
    primes[l]=number
  }
  return true;
}

func Primes_up_to(number int64) (results []int64)
{
	space := make([]bool,number);
	p := int64(2);
	for {
		fmt.Printf("%d\n", p);
		// mark all multiples of p non-prime
		for i := p+p; i<number; i+=p {
			space[i]=true;
		}
		if int64(math.Pow(float64(p),2)) > number {
			break
		}
		// find the next p
		for i:=p+1; i< number; i+=1 {
			if space[i]==false{
				p=i;
				break
			}
		}	
	}
	tresults := vector.New(0);
	for i:=int64(2); i<number; i++ {
		if space[i]==false {
			tresults.Push(i);
		}
	}
	results = make([]int64, tresults.Len());
	i := 0;
	for v := range tresults.Iter() {
		results[i] = v.(int64);
		i++
	}
	
	return;
}