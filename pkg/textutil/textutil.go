// Package textutil is the NESTED-library half of the monorepo demo: its own
// go.mod inside a repo that also ships CLIs. Its tag stream is the directory
// path — pkg/textutil/vX.Y.Z — because the module proxy resolves no other
// shape for a submodule.
package textutil

// Reverse returns s with its runes in reverse order.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
