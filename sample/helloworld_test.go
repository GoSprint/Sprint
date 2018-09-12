package main

import "testing"

func TestMain(t *testing . T) {
	t.Log("Pass")
}

func TestAdd(t *testing . T) {
	if (2 != add(1, 1)) {
		t.Error("Fail on 1 + 1 = 2")
	}


	if (5 != add(2, 3)) {
		t.Error("Fail on 2 + 3 = 5")
	}
}
