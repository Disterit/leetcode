package main

import "fmt"

func checkVoid(s string) string {
	var i int

	for _, v := range s {
		if v == ' ' {
			i += 1
		} else {
			break
		}
	}

	return s[i:]
}

func myAtoi(s string) int {

	var digit int
	sign := 1
	i := 0

	s = checkVoid(s)

	if len(s) == 0 {
		return 0
	}

	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		sign = 1
		i++
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit = digit*10 + int(s[i]-'0')

		if digit > 2147483647 && sign == 1 {
			digit = 2147483647
		} else if digit > 2147483648 && sign == -1 {
			return -2147483648
		}

		i++
	}

	return sign * digit
}

func main() {
	//fmt.Println(myAtoi("42"))
	//fmt.Println(myAtoi("-42"))
	//fmt.Println(myAtoi("1337c0d3"))
	//fmt.Println(myAtoi("0-1"))
	//fmt.Println(myAtoi("words and 987"))
	//fmt.Println(myAtoi("   -042"))
	//fmt.Println(myAtoi("-91283472332"))
	//fmt.Println(myAtoi("+1"))
	//fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("-91283472332"))
}
