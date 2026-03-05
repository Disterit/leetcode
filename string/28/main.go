package main

import "fmt"

func strStr(haystack string, needle string) int {

	res := -1

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return res
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
}

//runtime 100%
//memory 53.92%
