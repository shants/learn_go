package main

import (
	"fmt"
	"sort"
	"strconv"
)

type emp struct {
	name string
	age int
}

type empSeq []emp

func (s empSeq) Len() int {
return len(s)
}


func (s empSeq) Less(i, j int) bool {
	return s[i].age < s[j].age
}
func (s empSeq) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}



func main() {
e1 :=emp{name:"Shantanu",age:20}
e2 :=emp{name:"at15",age:15}
e3 := emp{name:"at25",age:25}
e4 := emp{name:"at251",age:251}
 s :=empSeq{e1,e2,e3,e4}
sort.Sort(s)
for i:=0; i < s.Len(); i++ {
	fmt.Println(strconv.Itoa(s[i].age) + " : " + s[i].name )
}
	}
