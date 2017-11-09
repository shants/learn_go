package main

import (
	"fmt"
	"sort"
	"strconv"
)

type num struct {
	idx int
	value int
}

type numSeq []num


func twoSum(nums []int, target int) []int {
	n :=  numSeq{}//make([]num, len(nums))
	ans := make([]int,2)
	for i:=0; i< len(nums); i++ {
		e1 := num{idx:i, value:nums[i]}
		n = append(n,e1 )
	}
	sort.Sort(n)
	k := 0
	j := len(nums)-1
	for k<j {
		if(n[k].value + n[j].value == target) {
			ans[0]= n[k].idx
			ans[1]=n[j].idx
			break
		}
		if (n[k].value + n[j].value < target){
			k++
		}else if (n[k].value + n[j].value > target) {
			j--
		}
	}

	return ans
}


func (s numSeq) Len() int {
	return len(s)
}


func (s numSeq) Less(i, j int) bool {
	return s[i].value < s[j].value
}
func (s numSeq) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}



func main() {
	fmt.Println("Start main")
	nums := []int{2,7,15,11 }
	ans := twoSum(nums,9)
	fmt.Println("after twofunc")
	fmt.Println(strconv.Itoa(ans[0]) + " : " + strconv.Itoa(ans[1]))

}
