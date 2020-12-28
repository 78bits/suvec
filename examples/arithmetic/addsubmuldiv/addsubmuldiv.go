package main

import (
	"fmt"

	"github.com/78bits/suvec"
)

func main() {
	fmt.Println("Test")

	a := suvec.Init([][]float64{{0.0, 0.1, 0.2}, {1.2, 1.3, 1.4}})

	fmt.Println(a)
}
