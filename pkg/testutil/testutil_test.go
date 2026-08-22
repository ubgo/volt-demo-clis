package testutil

import "testing"

func TestMustEqualPasses(t *testing.T) {
	MustEqual(t, 42, 42)
}
