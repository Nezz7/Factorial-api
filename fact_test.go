package main

import "testing"


func TestFactorial(t *testing.T){
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
	

}