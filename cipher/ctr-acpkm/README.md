## CTR-ACPKM
CTR-ACPKM is a cipher block mode with a limited key pressure as described in [GOST](./docs/gost-3412-2015.pdf).

#### Overview
Key pressure means that keys are automatically updated after configured 
amount of data is encrypted with the given key. The configuration should be
chosen with respect to expected attacks.

#### Implementation

Every key change creates a new `cipher.Block` instance 
to reset the state with a new key. Its key length must be divisible by
block size in order to do acpkm (generate new key). Thus only keys 16 and 32
are allowed for `crypto/aes`.

The block mode takes `BlockFactory` to be able to create the implementation of
`cipher.Block` with new key. Factories for [kuznyechik](https://github.com/pkositsyn/kuznyechik) 
and [aes](https://golang.org/src/crypto/aes/cipher.go#L28) ciphers are exposed for usage.

#### Tests

One test was taken from GOST and others were generated according to
expected equivalent behaviour.

### Benchmarks

```
// Kuznyechik encrypt is ~100 MB/s
BenchmarkKuznyechikCtrAcpkm-4  376836    3137 ns/op    81.60 MB/s    0 B/op    0 allocs/op
BenchmarkAesCtrAcpkm-4         1344072   880 ns/op     290.83 MB/s   0 B/op    0 allocs/op
```

