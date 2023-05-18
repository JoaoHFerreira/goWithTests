package array_and_slices

func Sum(numbers [5]int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}
