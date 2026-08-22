package textutil

import "testing"

func TestReverse(t *testing.T) {
	for in, want := range map[string]string{"volt": "tlov", "": "", "héllo": "olléh"} {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
		}
	}
}
