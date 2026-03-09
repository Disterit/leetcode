package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {

	var result []string

	stack := strings.Split(path, "/")

	for i := 0; i < len(stack); i++ {
		if stack[i] == "" || stack[i] == "." {
			continue
		}

		if stack[i] == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			continue
		}

		result = append(result, stack[i])
	}

	res := strings.Join(result, "/")

	return "/" + res
}

func main() {
	fmt.Println(simplifyPath("/home//foo/"))                      // "/home/foo"
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures")) // "/home/user/Pictures"
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))            // "/.../b/d"
	fmt.Println(simplifyPath("/../"))                             // /
	fmt.Println(simplifyPath("/a/./b/../../c/"))                  // "/c"

}

//runtime 100.00%
//memory 86.07%
