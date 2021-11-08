package main

import "testing"

func TestAdd(t *testing.T) {
	v := Add(2, 3)
	if v != 5 {
		t.Errorf("Add(5) failed. Go %v, expected 2,3.", v)
	}
}
