// Package testutil is the demo's INTERNAL package: tested by volt ci like
// everything else, but never released — see .volt.yml (`internal: true`).
package testutil

import "testing"

// MustEqual fails the test when got != want, with a uniform message.
func MustEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
