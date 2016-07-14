package main

import (
	"fmt"
)

var b = true
var aslice = []int{1, 2, 4, 8, 16, 32, 64, 128}
var amap = map[string]string{"hello": "world"}

func main() {
	//Basic, note that there are no ( )
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Leave pre and post statements empty
	for b == false {
		fmt.Println(b)
		b = true
	}

	//For is go's while, just leave off the ;
	for b == true {
		fmt.Println(b)
		b = false
	}

	//Infinite Loop
	for {
		break
	}

	//Iterate over slice or map, index is optional
	for index, value := range aslice {
		fmt.Println(index, " ", value)
	}

	for value := range aslice {
		fmt.Println(value)
	}

	for key, value := range amap {
		fmt.Println(key, " ", value)
	}
}
