package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {

	var res string

	var backtrack func(s string, n int) string

	backtrack = func(s string, n int) string {
		if n == 0 {
			return s
		}

		if s == "" {
			return backtrack("1", n-1)
		} else if s == "1" {
			return backtrack("11", n-1)
		}

		var fnS string
		count := 1

		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				count++
			} else {
				fnS += strconv.Itoa(count) + string(s[i])
				count = 1
			}
		}

		fnS += strconv.Itoa(count) + string(s[len(s)-1])

		return backtrack(fnS, n-1)
	}

	res = backtrack("", n)

	return res
}

func main() {
	fmt.Println(countAndSay(4)) // 1211
	fmt.Println(countAndSay(5)) // 111221
	fmt.Println(countAndSay(6)) // 312211
}

//runtime 41.06%
//memory 22.81%
