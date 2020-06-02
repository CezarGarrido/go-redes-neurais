package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sigmoid(10000.0))
}

func Sigmoid(x float64) float64 {
	return (1 / (1 + math.Exp(-x)))
}
