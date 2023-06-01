package jsonmap_test

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fatal()
	}
}

func assertEqual(t *testing.T, a, b any) {
	if a != b {
		t.Fatalf("%v != %v", a, b)
	}
}
