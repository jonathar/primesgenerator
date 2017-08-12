package main

import (
	"errors"
)

// NaiveSieve implements the sieve of Eratosthenes
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
type NaiveSieve struct {
	composites map[Number][]Number
	generator  CandidateGenerator
	sieve      func() Number
}

// NewNaiveSieve constructs a NaiveSieve
func NewNaiveSieve() *NaiveSieve {
	m := make(map[Number][]Number)
	return &NaiveSieve{generator: &naiveCandidateGenerator{current: 1}, composites: m}
}

// NextPrime returns the next prime in sequence
func (s *NaiveSieve) NextPrime() (Number, error) {
	for true {
		candidate := s.generator.NextCandidate()
		if _, present := s.composites[candidate]; !present {
			s.composites[candidate*candidate] = []Number{candidate}
			return candidate, nil
		}
		for _, c := range s.composites[candidate] {
			if _, present := s.composites[candidate+c]; !present {
				s.composites[candidate+c] = []Number{}
			}
			s.composites[candidate+c] = append(s.composites[candidate+c], c)
		}
		delete(s.composites, candidate)
	}
	return 0, errors.New("Unreachable! Should not execute this instruction!")
}

type naiveCandidateGenerator struct {
	current Number
}

func (g *naiveCandidateGenerator) NextCandidate() Number {
	current := g.current
	g.current++
	return current
}
