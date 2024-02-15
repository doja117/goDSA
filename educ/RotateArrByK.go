package educ

import "fmt"

func RotateByK(data []int, k int) {
	new_arr := data[:k]
	fmt.Println(new_arr)
	for i, _ := range data {
		if i+k < len(data) {
			data[i] = data[i+k]
		}
	}
	for i := 0; i < len(new_arr); i++ {
		data[len(data)-1-k+i] = new_arr[i]
	}

}
