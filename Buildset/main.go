package main

import (
	"fmt"
)

type set map[string]struct{}

func main() {
	x := make(set)
	x["a"] = struct{}{}
	x["b"] = struct{}{}
	fmt.Println(getvalue(x))
}

func getvalue(x set) []string {
	var ret []string
	for k, _ := range x {
		ret = append(ret, k)
	}
	return ret
}

/*Print "a" and "b"*/
