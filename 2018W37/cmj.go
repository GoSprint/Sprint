package main

func Fibonacci_cmj (input int) int {
	if (input == 0) {
		return 1;
	} else if (input == 1) {
		return 2;
	} else {
		return 3 * Fibonacci_cmj(input - 1) - 2 * Fibonacci_cmj(input - 2)
	}
}

