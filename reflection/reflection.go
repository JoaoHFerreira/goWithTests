package main

import "fmt"

func main() {
	result := numberAdd(1, 2)
	string_result := StringSum("yes or ", "no")
	fmt.Println(result)
	fmt.Println(string_result)

}

func numberAdd(number_one, number_two interface{}) interface{} {
	return number_one + number_two
}

func StringSum(string_one, string_two string) string {
	return string_one + string_two
}
