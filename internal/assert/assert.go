package assert

import (
	"strings"
	"testing"
)


func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}

/*
	// 14.5 Mocking dependencies: Testing the snippetView handler

	// To check that the request body contains some specific content, rather than being exactly equal to it.
	// Add a new StringContains() function to our assert package
*/
func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got %q; expected to contain: %q", actual, expectedSubstring)
	}
}