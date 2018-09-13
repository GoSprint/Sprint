package main

import (
	"bufio"
	"os"
)


func Fibonacci_cmj (input int) int {
	if (input == 0) {
		return 1;
	} else if (input == 1) {
		return 2;
	} else {
		return 3 * Fibonacci_cmj(input - 1) - 2 * Fibonacci_cmj(input - 2)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NameTimes(times int, name string, fd *bufio.Writer) string {
	return_string := name

	for i := 1; i < times; i++ {
		return_string = return_string + ", " + name
	}
	return return_string
}

func FileOperator_cmj(x int, y int, z int) string {
	filename := "sample"

	f, err := os.Create( filename );
	check(err)

	fd := bufio.NewWriter(f)
	tmp_string := ""
	tmp_string = NameTimes(x, "CMJ", fd)
	tmp_string = tmp_string + ", " + NameTimes(y, "AWON", fd)
	tmp_string = tmp_string + ", " + NameTimes(z, "JAY", fd)

	_, err = fd.WriteString(tmp_string)
	check(err)

	fd.Flush()
	f.Sync()
	f.Close()
	return filename
}

