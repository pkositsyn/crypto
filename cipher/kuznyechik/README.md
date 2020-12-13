## Kuznyechik cipher

"Kuznyechik" is a russian 128bit block cipher with 256bit key length, specified in [GOST](./docs/GOST.pdf) (russian standard).

## Implementation

The implementation of this cipher is done in the way, that it corresponds to `Block` interface of `"crypto/cipher"` standard package. 
This allows to use cipher with different block modes such as CBC, CFB, CTR, OFB and etc. It is well-optimized and tested - feel free to check that.

Internally, the implementation uses lookup tables and 64 bit xor, which allows the cipher to have a high throughput. 
These lookup tables are not precalculated but initialized on startup in order not to slow down the execution time.

## Benchmarks

See the table below. The last benchmark corresponds to initializing lookup tables.

```
BenchmarkNewCipher-4    943480     1170 ns/op       4377.52 MB/s   320 B/op   1 allocs/op
BenchmarkEncrypt-4      7613686    143 ns/op        111.82 MB/s    0 B/op     0 allocs/op
BenchmarkDecrypt-4      7739772    154 ns/op        103.74 MB/s    0 B/op     0 allocs/op
BenchmarkInitTables-4   675        1721889 ns/op                   1024 B/op  64 allocs/op
```
