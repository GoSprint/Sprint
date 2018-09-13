package main

import (
	"strings"
	"io/ioutil"
)

func Fibonacci_jay(data int) int {
	if (data == 0) {
		return 1;
	} else if (data == 1) {
		return 2;
	} else {
		return 3 * Fibonacci_jay(data - 1) - 2 * Fibonacci_jay(data - 2);
	}
}

func SetupArrayWithName(numName int, name string) []string {
	data := []string{}
	for i := 0; i < numName; i++ {
		data = append(data, name)
	}
	return data
}

func Name_jay(numCmj int, numAwon int, numJay int) string {
	data := []string{}
	data = append(data, SetupArrayWithName(numCmj, "CMJ")...)
	data = append(data, SetupArrayWithName(numAwon, "AWON")...)
	data = append(data, SetupArrayWithName(numJay, "JAY")...)
	str := strings.Join(data, ", ")
	databytes := []byte(str)
	err := ioutil.WriteFile("jay_filetest", databytes, 0644)
	if err != nil {
		panic(err)
	}
	return "jay_filetest"
}
