package main

import "fmt"

var maps = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	var output []string

	var backtrack func(combination string, nextDigits string)

	backtrack = func(combination string, nextDigits string) {
		if len(nextDigits) == 0 {
			output = append(output, combination)
			return
		}

		nextDigit := nextDigits[0]
		letters := maps[string(nextDigit)]

		for _, letter := range letters {
			backtrack(combination+string(letter), nextDigits[1:])
		}
	}

	backtrack("", digits)

	return output
}

func main() {
	fmt.Println(letterCombinations("23"))
}

//runtime 100%
//memory 55.32%
