// Minimum Element After Replacement With Digit Sum
package main

func minElement(nums []int) int {
	minVal := 10000

	for _, num := range nums {
		sum := 0

		for num > 0 {
			sum += num % 10
			num /= 10
		}

		if sum <  minVal{
			minVal = sum
		}
	}
	return minVal
}