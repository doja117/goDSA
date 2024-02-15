package educ

func LargestSumSubArray(data []int) int {
	ans := 0
	curr_sum := 0
	for _, v := range data {
		curr_sum += v
		if curr_sum < 0 {
			curr_sum = 0
		} else {

			ans = max(curr_sum, ans)
		}
	}
	return ans
}
