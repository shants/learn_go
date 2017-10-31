package main

import (
	"fmt"
	"math"
)



func main() {
	var n int
	var a, b float64
	fmt.Scan(&n)
	var values = make([][]float64, n)
	for j:=0; j < n; j++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		values[j] = make([]float64,2)
		values[j][0] = a
		values[j][1] = b
	}
	for i:=0; i < n ; i++ {
		fmt(calculator(values[i][0],values[i][1] ))
	}
}

func calculator(i,j float64 ) float64 {
	var a1, b1, c1 int
	var base1 = float64(10)

	var sum float64 = 0
	idx := 0
	for j>0 || i > 0 {
		a1 = int (math.Mod(i, base1))
		b1 = int (math.Mod(j, base1))
		i = math.Floor(i/10)
		j = math.Floor(j /10)
		c1 = a1 + b1
		c1 = c1 %10
		sum = (math.Pow10(idx) * float64(c1)) + sum

		idx++
	}
	return sum
}