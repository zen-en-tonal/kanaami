package host

import "testing"

func TestMatch(t *testing.T) {
	if !Host("github.com").Match(Host("github.com")) {
		t.Error("same domain must match")
	}
}
