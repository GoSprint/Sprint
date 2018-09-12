package main

import (
	"fmt"
	"testing"
)

func Test (t *testing . T) {
	f := []func(int) int {
		Fibonacci_cmj,
	}
	cases := [][]int {
		{0, 1},
		{1, 2},
	}

	for i := 0; i < len(f); i++ {
		for j := 0; j < len(cases); j++ {
			if (cases[j][1] != f[i](cases[j][0])) {
				fmt.Printf("F[%d] = %d <> %d\n", cases[j][0], cases[j][1], f[i](cases[j][0]))
				t.Fail()
			}
		}
	}
}
