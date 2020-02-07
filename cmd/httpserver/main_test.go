package main

import "testing"

func TestSomething(t *testing.T) {
	if 2 != 2 {
		t.Errorf("Catastrophic error")
	}
}
