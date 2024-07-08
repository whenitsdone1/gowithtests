package sum

func Sum(a []int) int {
	var total int
	for _, val := range a {
		total += val
	}
	return total
}

func SumTails(numsToSum ...[]int) []int {
	var sums []int

	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
