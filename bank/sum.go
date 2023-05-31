package bank

func Sum(nums []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(nums, add, 0)
}

func SumAllTails(nums ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(nums, sumTail, []int{})
}
