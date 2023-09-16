package main

import "fmt"

func main() {
	nums := []int{4, 2, 5, 7}

	lastEven := 0
	lastOdd := 1
	result := make([]int, len(nums))
	for _, num := range nums {
		if num%2 == 0 && lastEven < len(result) {
			result[lastEven] = num
			lastEven += 2
			continue
		}

		if num%2 == 1 && lastOdd < len(result) {
			result[lastOdd] = num
			lastOdd += 2
		}
	}

	fmt.Println(result)

}
