package main

import "fmt"

func generateParenthesis(n int) []string {

	output := []string{}

	var backtrack func(combinations string, open, close int)

	backtrack = func(combinations string, open, close int) {
		if len(combinations) == 2*n {
			output = append(output, combinations)
			return
		}

		if open < n {
			backtrack(combinations+"(", open+1, close)
		}

		if close < open {
			backtrack(combinations+")", open, close+1)
		}

	}

	backtrack("", 0, 0)

	return output
}

func main() {
	fmt.Println(generateParenthesis(3)) //["((()))","(()())","(())()","()(())","()()()"]
}

//runtime 100%
//memory 80.12%
