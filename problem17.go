// Problem 16 - http://projecteuler.net/
//
// If all the numbers from 1 to 1000 (one thousand) 
// inclusive were written out in words, how many letters would be used?

package main

import (
	"fmt";
	"strings";
	"container/vector";
)

func number_to_words(num int) (s string)
{
	components := vector.NewStringVector(0);
	var ones = map[int] string {
		0: "", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
		6: "six", 7: "seven", 8: "eight", 9: "nine",
	};
	var teens = map[int] string {
		0: "ten", 1: "eleven", 2: "twelve", 3: "thirteen", 4: "fourteen", 5: "fifteen",
		6: "sixteen", 7: "seventeen", 8: "eighteen", 9: "nineteen",
	};
	var tens = map[int] string {
		2: "twenty", 3: "thirty", 4: "forty", 5: "fifty",
		6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety",
	};
	if num==1000 {
		return "one thousand"
	}
	t := 0;
	if num/100!=0 {
		t=num/100;
		components.Push(ones[t]);
		components.Push("hundred")
	}
	num=num%100;
	if num!=0 && components.Len()>0 { 
		components.Push("and");
		}
	if num/10!=0 {
		t=num/10;
		if t>1 {
			components.Push(tens[t]);
		}
		else {
			if t==1 {
				components.Push(teens[num%10])
			}
		}
	}
	if num<10 || num >19 {
		components.Push(ones[num%10]);
	}
	
	return strings.Join(components.Data(), " ");
}

func main() {
	count := 0;
	lens := 0;
	for i:=1; i<=1000; i++ {
		lens = len(strings.Join(strings.Split(number_to_words(i), " ",0), ""));
		fmt.Printf("%s (%d)\n", number_to_words(i), lens);
		count+= lens;
	}
	fmt.Printf("%d\n", count);
}
