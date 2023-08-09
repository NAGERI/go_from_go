package main

func SumAll(numbersToSum ...[]int) []int {

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	// make - creates a slice with capacity of len of numbersToSum

	for i, number := range numbersToSum {
		sums[i] = Sum(number)
	}
	return sums

}
