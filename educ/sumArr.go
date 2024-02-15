package educ

func SumArr(data []int) int {
	ans := 0

	for _, v := range data {
		ans += v
	}
	return ans
}
