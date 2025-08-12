package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}

	// Loop through array
	for i, v := range arr {
		fmt.Printf("Index: %d Value: %d\n", i, v)
	}

	// Length and capacity
	fmt.Println("Array length:", len(arr))
	fmt.Println("Array capacity:", cap(arr))

	// Elements
	fmt.Println("Array elements:", arr)

	// Sum
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Array sum:", sum)

	// Average
	avg := float64(sum) / float64(len(arr))
	fmt.Println("Array average:", avg)

	// Max
	fmt.Println("Array max:", max(arr))
}

// max returns the largest element in an int slice
func max(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}
	m := nums[0]
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}
