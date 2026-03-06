package main

import "fmt"

func addBinary(a string, b string) string {

	if len(a) < len(b) {
		a, b = b, a
	}

	var res string
	var degree uint8
	len1, len2 := len(a), len(b)

	for i := 0; i < len1-len2; i++ {
		b = "0" + b
	}

	for i := 0; i < len1; i++ {

		num1 := a[len1-i-1] - '0'
		num2 := b[len1-i-1] - '0'

		if num1+num2+degree == 3 {
			res = "1" + res
			degree = 1
		} else if num1+num2 == 2 {
			res = "0" + res
			degree = 1
		} else if num1+num2+degree == 2 {
			res = "0" + res
			degree = 1
		} else if num1+num2+degree == 1 {
			res = "1" + res
			degree = 0
		} else {
			res = "0" + res
		}
	}

	if degree == 1 {
		res = "1" + res
	}

	return res
}

func main() {
	fmt.Println(addBinary("1111", "11"))
}

// 111
// 011

// 1
