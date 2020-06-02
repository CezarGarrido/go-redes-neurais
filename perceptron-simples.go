package main

import (
	"fmt"
)

func main() {
	var entradas = []int{1, 7, 5}
	var pesos = []float64{0.8, 0.1, 0}

	fmt.Println(stepFunction(soma(entradas, pesos)))
}

func soma(entradas []int, pesos []float64) float64 {
	var s float64 = 0

	for _, e := range entradas {
		for _, p := range pesos {
			s += float64(e) * p
		}
	}
	return s
}

func stepFunction(soma float64) int {
	if soma >= 1.0 {
		return 1
	}
	return 0
}
