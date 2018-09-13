package main

func Fibonacci_jay(data int) int {
	if (data == 0) {
		return 1;
	} else if (data == 1) {
		return 2;
	} else {
		return 3 * Fibonacci_jay(data - 1) - 2 * Fibonacci_jay(data - 2);
	}
}

func FindCMJ_jay() int {
}
