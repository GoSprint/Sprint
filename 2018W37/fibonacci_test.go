package main

import (
	"fmt"
	"testing"
	"io/ioutil"
	"strings"
	"os"
)

func Test_Fibonacci (t *testing . T) {
	f := []func(int) int {
		Fibonacci_cmj,
		Fibonacci_jay,
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

func GetNameNum(nameArray []string, name string) int {
	number := 0;
	for i := 0; i < len(nameArray); i++ {
		if name == nameArray[i] {
			number++
		}
	}
	return number;
}

func CheckFileOperatorNumber(nameArray []string, numArray []int) bool {
	checkArray := []int{GetNameNum(nameArray, "CMJ"),
						GetNameNum(nameArray, "AWON"),
						GetNameNum(nameArray, "JAY")}

	if len(nameArray) != checkArray[0] + checkArray[1] + checkArray[2] {
		return false
	}

	for i := 0; i < len(checkArray); i++ {
		if checkArray[i] != numArray[i] {
			return false;
		}
	}
	return true
}

func Test_Fileoperator (t *testing . T) {
	f := []func(int, int, int) string {
		Name_jay,
	}
	cases := [][]int {
		{3, 2, 1},
		{5, 7, 9},
		{7, 2, 3},
	}

	for i := 0; i < len(f); i++ {
		for j := 0; j < len(cases); j++ {
			filename := f[i](cases[j][0], cases[j][1], cases[j][2]);
			b, err := ioutil.ReadFile(filename);
			if err != nil {
				t.Fail()
			}
			array := strings.Split(strings.TrimSpace(string(b)), ", ")
			os.Remove(filename)
			if !CheckFileOperatorNumber(array, cases[j]) {
				t.Fail()
			}
		}
	}
}
