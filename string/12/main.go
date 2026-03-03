package main

import (
	"fmt"
	"strings"
)

//var maps = map[string]int{
//	"M":  1000,
//	"CM": 900,
//	"D":  500,
//	"CD": 400,
//	"C":  100,
//	"XC": 90,
//	"L":  50,
//	"XL": 40,
//	"X":  10,
//	"IX": 9,
//	"V":  5,
//	"IV": 4,
//	"I":  1,
//}

var maps = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var arrays = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var i int
	var builder strings.Builder

	for num > 0 && i < len(arrays) {
		if num >= arrays[i] {
			num -= arrays[i]
			builder.WriteString(maps[arrays[i]])
		} else {
			i++
		}
	}

	return builder.String()
}

func main() {
	fmt.Println(intToRoman(3749)) //MMMDCCXLIX
	fmt.Println(intToRoman(58))   //LVIII
	fmt.Println(intToRoman(1994)) //MCMXCIV
}

//runtime 100%
//memory 43.21%
