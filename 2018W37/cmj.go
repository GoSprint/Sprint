package main

import "fmt"

func main() {
	fmt.Printf("F[10] = %d\n", Fibonacci_cmj(10))

	cases := [][]int {
		{0, 1},
		{1, 2},
	}

	for i := 0; i < len(cases); i++ {
		fmt.Printf("%d %d\n", cases[i][0], cases[i][1])
	}
}

func Fibonacci_cmj (input int) int {
	if (input == 0) {
		return 1;
	} else if (input == 1) {
		return 2;
	} else {
		return 3 * Fibonacci_cmj(input - 1) - 2 * Fibonacci_cmj(input - 2)
	}
}

