package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fact(1050)
	}
}
func TestFactorial(t *testing.T) {
	res := fact(0)
	if res != 1 {
		t.Errorf("fact(0) = %d expected 1", res)
	}

	res = fact(1)
	if res != 1 {
		t.Errorf("fact(1) = %d expected 1", res)
	}

	res = fact(5)
	if res != 120 {
		t.Errorf("fact(5) = %d expected 120", res)
	}

	res = fact(-120)
	if res != -1 {
		t.Errorf("fact(-120) = %d expected -1", res)
	}

}
