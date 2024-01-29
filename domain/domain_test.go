package domain

import "testing"

func TestMatch(t *testing.T) {
	if !Domain("github.com").Match(Domain("github.com")) {
		t.Error("same domain must match")
	}
}
