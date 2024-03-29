// +build ignore

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		fmt.Println(sqrt(-x) + "i")
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
