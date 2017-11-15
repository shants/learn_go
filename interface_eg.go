package main

import (
	"fmt"
	"math"

	)

type vertex struct {
	x float64
	y float64
}

// Abser interface implements Abs function
// go has no explicit implements
type Abser interface {
	Abs() float64
}

// here type v implements Abs function hence Abser interface
func (v *vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Stringer interface has String method
type person struct {
	name string
	age int
}

// type person implements Stringer

func (p *person) String() string {
	return fmt.Sprintf("%v %v years", p.name, p.age)
}

func (*person) PrintHello() string {
	return fmt.Sprintf("Hello")
}

func (p *person) Abs() float64 {
	return float64(p.age)
}

type IPAddr [4]byte
// y by value and not pointer
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])

}

func testEmptyInter(a interface{}){

	fmt.Println("Hello empty")
}

func implAbs(abser Abser) float64 {
	return abser.Abs()
}

func main() {
 v:=vertex{3,4}
 fmt.Println(v.Abs())

	a := person{"Arthur Dent", 42}
	fmt.Println(a)
	fmt.Println(a.PrintHello())
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	testEmptyInter(v)
	testEmptyInter(a)

	fmt.Println(implAbs(&v))
	fmt.Println(implAbs(&a))
 	}

