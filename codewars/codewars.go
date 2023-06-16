package main

func Between(a, b int) []int {
	var between_number_list = []int{}
	for i := a; i <= b; i++ {
		between_number_list = append(between_number_list, i)
	}
	return between_number_list
}
