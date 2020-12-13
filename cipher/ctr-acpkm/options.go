package ctracpkm

// CTROptions are additional options for ctr-acpkm
type CTROptions int

// WithGammaSize set gamma size.
// Gamma is the length of a block to be xor'ed with next ctr secret value (actually with gamma most significant bytes of it).
// Block size must be divisible by gamma size.
// Default is block size. Processing is slowed by (block size / gamma size) times.
func WithGammaSize(size int) CTROptions {
	return CTROptions(size)
}
