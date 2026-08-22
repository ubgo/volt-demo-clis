// Package mathx is a PUBLISHED library in the monorepo demo — released and
// proxy-resolvable, coexisting with the internal pkg/testutil next to it.
package mathx

// Clamp bounds v to [lo, hi]. lo > hi is a caller bug and panics rather than
// silently swapping.
func Clamp(v, lo, hi int) int {
	if lo > hi {
		panic("mathx.Clamp: lo > hi")
	}
	switch {
	case v < lo:
		return lo
	case v > hi:
		return hi
	default:
		return v
	}
}
