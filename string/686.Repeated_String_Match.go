package main

import "fmt"

func main686() {
	A := "abcd"
	B := "cdabcdab"
	ret := repeatedStringMatch(A, B)
	fmt.Println(ret)
}

func repeatedStringMatch(A string, B string) int {
	str := A
	num := 0
	for i, _ := range B {
		num += int(B[i] - 'a')
	}
	for len(str) < len(B)+2*len(A) {
		if subString(str, B, num) {
			return len(str) / len(A)
		}
		str += A
	}
	return -1
}

func subString(longStr, shortStr string, num int) bool {
	if len(longStr) < len(shortStr) {
		return false
	}
	n := len(shortStr)
	temp := 0
	for i := 0; i < n; i++ {
		temp += int(longStr[i] - 'a')
	}
	if temp == num && longStr[:n] == shortStr {
		return true
	}
	for i := 0; i+n < len(longStr); i++ {
		j := i + n
		temp -= int(longStr[i] - 'a')
		temp += int(longStr[j] - 'a')
		if temp == num && longStr[i+1:j+1] == shortStr {
			return true
		}
	}
	return false
}
