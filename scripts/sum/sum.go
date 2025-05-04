package main

func Sum(s []int) int {
	var sum int
	for _, number := range s {
		sum += number
	}

	return sum
}

func SumAll(s ...[]int) []int {
	var result []int
	for _, slice := range s {
		result = append(result, Sum(slice))
	}
	return result
}

func SumAll2(s ...[]int) []int {
	result := make([]int, len(s))
	for i, slice := range s {
		result[i] = Sum(slice)
	}
	return result

}

func SumAllTails(s ...[]int) []int {
	result := make([]int, len(s))
	for i, slice := range s {
		if len(slice) == 0 {
			result[i] = 0
			continue
		}
		result[i] = Sum(slice[1:])
	}
	return result
}
