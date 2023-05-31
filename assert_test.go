package jsonmap_test

import "testing"

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}
