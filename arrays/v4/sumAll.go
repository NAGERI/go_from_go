package main

func SumAll(numbersToSum ...[]int) []int {

	var sum []int
	for _, number := range numbersToSum {
		sum = append(sum, Sum(number))
	}
	return sum

}
