package main

import (
	"testing"
	"io/ioutil"
	"strings"
	"os"
)

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
		FileOperator_cmj,
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
