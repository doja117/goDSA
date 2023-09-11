package main

import "fmt"

var ans [][]int

func returnArr(p []int, q []int) {
	if len(p) == 0 {
		ans = append(ans, q)
		return
	}
	returnArr(p[1:], q)
	q = append(q, p[0])
	returnArr(p[1:], q)

}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{}
	returnArr(arr1, arr2)
	for _, v := range ans {
		fmt.Println(v)
	}
}
