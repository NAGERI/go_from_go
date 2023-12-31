package main

func SumAllTails(numberToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
