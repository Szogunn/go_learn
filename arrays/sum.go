package arrays

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTailes(slices ...[]int) []int {
	result := make([]int, len(slices))

	for i, slice := range slices {
		if len(slice) == 0 {
			result[i] = 0
		} else {
			result[i] = Sum(slice[1:])
		}
	}

	return result
}
