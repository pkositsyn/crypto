## Streebog hash function

Streebog is a russian standard cryptographic hash function 
described in [GOST](./docs/gost-34.11-2012.pdf). 
There are 2 versions of the hash - 256 and 512 bit.
This implementation is based on lookup table and thus well-optimized (see benchmarks).

### Note

Hash function starts encoding from the end. Thus, it cannot have state and
hash sum must be calculated entirely every time `Sum()` is called.
Moreover, marshalling hasher, which only state is consumed data, is useless.

Its memory consumption is linear to provided data until `Reset()` is called.
Usually reset is called after each *message* unit processing.

### Tests
Most of the tests for the cipher are taken from GOST. The other tests check
interface contracts for hash - [hash.Hash](https://pkg.go.dev/hash#Hash).

Test coverage for initial version is 100%

### Benchmarks

```
// 512 bit hash. Only Sum() function is heavy - it does most of work
BenchmarkSum512-4  14275  84178 ns/op  76.03 MB/s  64 B/op  1 allocs/op

// 256 bit hash
BenchmarkSum256-4  13179  83928 ns/op  76.26 MB/s  32 B/op  1 allocs/op

// Lookup table is initialized once with the first New... call
// BenchmarkInitLPSTable-4  102705  11036 ns/op  185.57 MB/s  0 B/op  0 allocs/op
```
