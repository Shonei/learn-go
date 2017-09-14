package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(cashUnits(1972.78))
}

func cashUnits(f float64) map[float64]int {
	arr := []float64{500.0, 200.0, 100.0, 50.0, 20.0, 10.0, 5.0, 2.0, 1.0, 0.5, 0.2, 0.1, 0.05, 0.02, 0.01}
	m := map[float64]int{}
	for _, v := range arr {
		m[v] = int(f / v)
		f = toFixed(f-float64(m[v])*v, 2)
	}

	return m
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
