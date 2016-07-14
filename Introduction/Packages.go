package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//Exported Function
	fmt.Println("My favorite number is", rand.Intn(10))
	//Exported Const
	fmt.Println(math.Pi)
}
