package main

// Number is a type alias for candidates and primes
type Number int

// Sieve generates primes
type Sieve interface {
	NextPrime() (Number, error)
}

// CandidateGenerator supplies integers to a sieve to be considered prime
type CandidateGenerator interface {
	NextCandidate() Number
}
