package main

import (
	"fmt"
)

//func sum(num1 string, num2 string) string {
//	if len(num2) > len(num1) {
//		num1, num2 = num2, num1
//	}
//
//	var res string
//	var degree uint8
//
//	for i := 0; i < len(num1); i++ {
//		var num uint8
//		if i < len(num2) {
//			num = (num1[len(num1)-i-1] - '0') + (num2[len(num2)-i-1] - '0') + degree
//		} else {
//			num = num1[len(num1)-i-1] - '0' + degree
//		}
//
//		degree = 0
//
//		if num >= 10 {
//			degree = 1
//			num -= 10
//		}
//
//		res = string(num+'0') + res
//	}
//
//	if degree == 1 {
//		res = "1" + res
//	}
//
//	return res
//}
//
//func multiply(num1 string, num2 string) string {
//
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//
//	var res string
//	var degree string
//
//	for i := 0; i < len(num2); i++ {
//		if i > 0 {
//			degree += "0"
//		}
//
//		number := num2[len(num2)-1-i] - '0'
//		var j uint8
//		var sums string
//
//		for j = 0; j < number; j++ {
//			sums = sum(sums, num1)
//		}
//
//		res = sum(res, sums+degree)
//	}
//
//	return res
//}

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	len1, len2 := len(num1), len(num2)

	arr := make([]uint8, len1+len2)

	for i := 0; i < len1; i++ {
		number1 := num1[len1-i-1] - '0'
		for j := 0; j < len2; j++ {
			number2 := num2[len2-j-1] - '0'

			pos1 := i + j
			pos2 := i + j + 1

			degree := arr[pos1]

			res := number1*number2 + degree

			arr[pos1] = res % 10
			arr[pos2] += res / 10
		}
	}

	var res string
	lenArr := len(arr)

	if arr[lenArr-1] == 0 {
		lenArr--
	}

	for i := 0; i < lenArr; i++ {
		res += string(arr[lenArr-i-1] + '0')
	}

	return res
}

func main() {
	//fmt.Println(multiply("123", "456")) //56088
	fmt.Println(multiply("2", "3")) // 6
	fmt.Println(multiply("55", "55"))
}

//runtime 53.47%
//memory 42.08%
