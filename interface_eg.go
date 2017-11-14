package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x float64
	y float64
}

type Abser interface {
	Abs() float64
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
 v:=vertex{3,4}
 fmt.Println(v.Abs())
}

