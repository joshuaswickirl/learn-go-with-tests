package arraysandslices

func Sum(n []int) int {
	var total int
	for _, i := range n {
		total += i
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	sum := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
//     var sums []int
//     for _, numbers := range numbersToSum {
//         sums = append(sums, Sum(numbers))
//     }
//     return sums
// }

// SumAllTails calculates the sum of all items except for the head
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
