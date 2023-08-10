package main

func SumAllTails(numberToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numberToSum {
		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}
