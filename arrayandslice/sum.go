package arrayandslice

func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
