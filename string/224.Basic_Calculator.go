package main

import (
	"fmt"
)

func main() {
	formular := "(1+(4+5+2)-3)+(6+ 8)"
	//formular := "1 + 1"
	//formular := "2147483647"
	ret := calculate(formular)
	fmt.Println(ret)
}

func calculate(s string) int {
	stack := make([]int, 0)
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res = res + sign*num
			i--
		} else {
			switch s[i] {
			case '+':
				sign = 1
			case '-':
				sign = -1
			case '(':
				stack = append(stack, res, sign)
				res, sign = 0, 1
			case ')':
				sign := stack[len(stack)-1]
				num := stack[len(stack)-2]
				res = num + sign*res
				stack = stack[:len(stack)-2]
			}
		}
	}
	return res
}
