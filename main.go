package main

import (
	"fmt"
)

func main() {
	sieve := NewNaiveSieve()
	prime := Number(1)
	for i := 0; i <= 1000; i++ {
		prime, _ = sieve.NextPrime()
	}
	fmt.Printf("hello %d\n", prime)
}
