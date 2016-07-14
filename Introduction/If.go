package main

import (
	"fmt"
)

func main() {
	b1 := true
	//Basic if, note that there are no ( ), and { } are required
	if b1 {
		fmt.Println(b1)

	}

	//Can also execute a statement, scoped to the if & else
	if b2 := !b1; b2 {
		fmt.Println(b2)
	} else {
		fmt.Println(b2)
	}
}
