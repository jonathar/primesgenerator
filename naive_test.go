package main

import (
	"reflect"
	"testing"
)

func buildTestSieves() []Sieve {
	return []Sieve{NewNaiveSieve()}
}

func TestFirstPrimes(t *testing.T) {
	var firstprimestest = []Number{1, 2, 3, 5, 7, 11}
	for _, sieve := range buildTestSieves() {
		for _, tt := range firstprimestest {
			np, err := sieve.NextPrime()
			if err != nil {
				t.Errorf("Something went horribly wrong with %s. NextPrime returned ERROR: %s", reflect.TypeOf(sieve), err)
			}
			if np != tt {
				t.Errorf("Expected %d to be the next prime. %s.NextPrime() gave %d instead.", tt, reflect.TypeOf(sieve), np)
			}
		}
	}
}

func TestThousandthPrime(t *testing.T) {
	for _, sieve := range buildTestSieves() {
		for i := 0; i < 1000; i++ {
			_, err := sieve.NextPrime()
			if err != nil {
				t.Errorf("Something went horribly wrong. NextPrime returned ERROR: %s", err)
			}
		}
		if p, _ := sieve.NextPrime(); p != 7919 {
			t.Errorf("Expected 7,919 to be the 1,000th prime. Got %d instead.", p)
		}
	}
}
