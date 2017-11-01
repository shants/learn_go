package main

import (
	"fmt"
	"math"
)



func main() {
	var n int
	var a, b uint64
	fmt.Scan(&n)
	var values = make([][]uint64, n)
	for j:=0; j < n; j++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		values[j] = make([]uint64,2)
		values[j][0] = a
		values[j][1] = b
	}
	for i:=0; i < n ; i++ {
		fmt.Println(calculator(values[i][0],values[i][1] ))
	}
}

func calculator(i,j uint64 ) uint64 {
	var a1, b1, c1 int


	var sum uint64 = 0
	idx := 0
	for j>0 || i > 0 {

		a1 = int (i % 10)
		b1 = int (j % 10)
		i =  i /10
		j = j /10
		c1 = a1 + b1
		c1 = c1 %10
		sum = uint64(math.Pow10(idx) * float64(c1)) + sum

		idx++
	}
	return sum
}