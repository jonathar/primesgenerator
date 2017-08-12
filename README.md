# Primes Generator

Prime number generator written in golang. Experimenting with a few different
algorithms.

## Generation Algorithms

  * Sieve of Eratosthenes
  * (TODO) Narrow wheel implementation of Sieve - skipping multiples of 2 & 3
  * (TODO) Wide wheel implementation of Sieve - skipping multiples of 2, 3 & 5

 See this [paper](http://www.cs.hmc.edu/~oneill/papers/Sieve-JFP.pdf) for the optimizations

## Benchmarking the different implementations

```bash
  $> go test -bench=.
```
